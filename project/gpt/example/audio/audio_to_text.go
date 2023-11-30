package audio

import (
	"context"
	"fmt"
	"testing"

	openai "github.com/sashabaranov/go-openai"
)

/*
*
语音转文字
*/
func TestAudio(t *testing.T) {
	c := openai.NewClient("")
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: "ai.mp4",
	}
	resp, err := c.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return
	}
	fmt.Println(resp.Text)
}
