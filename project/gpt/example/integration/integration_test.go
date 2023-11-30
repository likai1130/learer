package integration

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

var c *openai.Client

func init() {
	c = openai.NewClient("")
}

/*
*

	做一个语音转文字，文字转语音的示例。

1. 语音输入问题
2. 语音转文字
3. 文生文
4. 文字转语音
5. 输出为MP3
*/
func TestIntegration(t *testing.T) {
	sourceAIfilePath := "./demo.mp3"
	resultAIfilePath := "result_ai.mp3"
	demoText := "请你用哲学家的角度，谈论一下苏格拉底的一生是怎么样的，要求客观一点!"

	// 文生音
	if err := textToAudio(demoText, sourceAIfilePath); err != nil {
		panic(err)
	}

	// 音生文
	text, err := audioToText(sourceAIfilePath)
	if err != nil {
		panic(err)
	}

	// 文生文
	toText, err := textToText(text)
	if err != nil {
		panic(err)
	}

	// 文生音
	if err = textToAudio(toText, resultAIfilePath); err != nil {
		panic(err)
	}

}

func audioToText(filePath string) (string, error) {
	ctx := context.Background()
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: filePath, //音频路径
	}
	resp, err := c.CreateTranscription(ctx, req)
	if err != nil {
		log.Fatalf("Transcription error: %v\n", err)
		return "", errors.WithMessage(err, "Transcription")
	}
	return resp.Text, nil
}

func textToAudio(text string, filePath string) error {
	ctx := context.Background()
	req := openai.CreateSpeechRequest{
		Model: openai.TTsModel1HD,
		Input: text,
		Voice: openai.VoiceAlloy,
		//ResponseFormat: "",
		//默认响应格式为“mp3”，但也可以使用“opus”、“aac”或“flac”等其他格式。
		//Opus：用于互联网流媒体和通信，低延迟。AAC：用于数字音频压缩，
		//YouTube、Android、iOS 首选。FLAC：用于无损音频压缩，受到音频爱好者存档的青睐。
		//Speed: 0,
	}

	response, err := c.CreateSpeech(ctx, req)
	if err != nil {
		return errors.WithMessage(err, "CreateSpeech")
	}
	defer response.Close()

	buf, err := io.ReadAll(response)
	if err != nil {
		return errors.WithMessage(err, "IOReadAll")
	}

	// save buf to file as mp3
	err = os.WriteFile(filePath, buf, 0644)
	if err != nil {
		return errors.WithMessage(err, "WriteFile")
	}
	return nil
}

const encoding = "cl100k_base"

// getTokenByEncoding
func getTokenByEncoding(text string, encoding string) (int, error) {
	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		return 0, errors.WithMessage(err, "GetEncoding")
	}
	token := tke.Encode(text, nil, nil)
	return len(token), nil
}

func textToText(text string) (string, error) {
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo1106,
		MaxTokens: 4096,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: text,
			},
		},
		Stream: true,
	}

	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return "", errors.WithMessage(err, "ChatCompletionStream")
	}
	defer stream.Close()

	builder := strings.Builder{}
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			t := builder.String()
			fmt.Println(builder.String())
			return t, nil
		}
		if err != nil {
			return "", errors.WithMessage(err, "Stream")
		}
		builder.WriteString(response.Choices[0].Delta.Content)
	}
}
