package assistants

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"testing"
)

// 创建线程
func createThread(ctx context.Context) (openai.Thread, error) {
	return client.CreateThread(ctx, openai.ThreadRequest{
		Messages: []openai.ThreadMessage{
			{
				Role:    openai.ThreadMessageRoleUser,
				Content: "Hello, World!",
				//FileIDs:  nil,
				//Metadata: nil,
			},
		},
	})
}

// 查询线程
func getThread(ctx context.Context, threadId string) {
	response, err := client.RetrieveThread(ctx, threadId)
	if err != nil {
		panic(err)
	}
	objToJson(response, "getThread")
}

// 删除线程
func deleteThread(ctx context.Context, threadId string) {
	response, err := client.DeleteThread(ctx, threadId)
	if err != nil {
		panic(err)
	}
	objToJson(response, "deleteThread")
}

func TestCRUDThread(t *testing.T) {
	ctx := context.Background()
	thread, err := createThread(ctx)
	if err != nil {
		panic(err)
	}
	threadId := thread.ID
	// 查询
	getThread(ctx, threadId)

	// 删除
	deleteThread(ctx, threadId)
}

func TestThread(t *testing.T) {
	c := openai.NewClient("")
	ctx := context.Background()
	response, err := c.CreateThread(ctx, openai.ThreadRequest{
		Messages: []openai.ThreadMessage{
			{
				Role:    openai.ThreadMessageRoleUser,
				Content: "Hello, World!", //第一次可以不写
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
