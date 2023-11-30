package speech

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
	"testing"
)

var client *openai.Client

func init() {
	client = openai.NewClient("")
}

/*
*
文字转语音
*/
func TestSpeechAI(t *testing.T) {
	ctx := context.Background()

	req := openai.CreateSpeechRequest{
		Model: openai.TTSModel1,
		Input: "问你一个问题，你知道为什么天空是蓝色的？请用AI回答我",
		Voice: openai.VoiceAlloy,
		//ResponseFormat: "",
		//默认响应格式为“mp3”，但也可以使用“opus”、“aac”或“flac”等其他格式。
		//Opus：用于互联网流媒体和通信，低延迟。AAC：用于数字音频压缩，
		//YouTube、Android、iOS 首选。FLAC：用于无损音频压缩，受到音频爱好者存档的青睐。
		//Speed: 0,
	}

	response, err := client.CreateSpeech(ctx, req)
	if err != nil {
		log.Fatalf("创建speech错误，err=%v", err)
	}
	defer response.Close()

	buf, err := io.ReadAll(response)
	if err != nil {
		log.Fatalf("ReadAll error ，err=%v", err)
	}

	// save buf to file as mp3
	err = os.WriteFile("test.mp3", buf, 0644)
	if err != nil {
		log.Fatalf("WriteFile error ，err=%v", err)
	}
}
