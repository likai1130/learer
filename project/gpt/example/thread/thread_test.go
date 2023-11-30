package thread

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"testing"
)

func TestThread(t *testing.T) {
	c := openai.NewClient("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")
	ctx := context.Background()
	response, err := c.CreateThread(ctx, openai.ThreadRequest{
		Messages: []openai.ThreadMessage{
			{
				Role:    openai.ThreadMessageRoleUser,
				Content: "Hello, World!",
				//FileIDs:  nil,
				//Metadata: nil,
			},
		},
		Metadata: nil,
	})
	if err != nil {
		panic(err)
	}

	marshal, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}

func TestThreadAndRun(t *testing.T) {
	assistantID := "asst_abc123"
	/*threadID := "thread_abc123"
	runID := "run_abc123"
	stepID := "step_abc123"
	limit := 20
	order := "desc"
	after := "asst_abc122"
	before := "asst_abc124"*/

	c := openai.NewClient("")
	ctx := context.Background()
	response, err := c.CreateThreadAndRun(ctx, openai.CreateThreadAndRunRequest{
		RunRequest: openai.RunRequest{
			AssistantID:  assistantID,
			Model:        nil,
			Instructions: nil,
			Tools:        nil,
			Metadata:     nil,
		},
		Thread: openai.ThreadRequest{
			Messages: []openai.ThreadMessage{
				{
					Role:    openai.ThreadMessageRoleUser,
					Content: "Hello, World!",
					//FileIDs:  nil,
					//Metadata: nil,
				},
			},
			//Metadata: nil,
		},
	})
	if err != nil {
		panic(err)
	}

	marshal, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))

}
