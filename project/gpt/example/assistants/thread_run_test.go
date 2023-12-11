package assistants

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"testing"
	"time"
)

// 创建run线程并且运行他
// createThreadRun = createMessage + createThread + createRun
func createThreadRun(ctx context.Context, assistantId string, msg string) (openai.Run, error) {
	return client.CreateThreadAndRun(ctx, openai.CreateThreadAndRunRequest{
		RunRequest: openai.RunRequest{
			AssistantID: assistantId,
			//Instructions: assistant.Instructions,
			//Tools:        assistant.Tools,
			//Metadata:     assistant.Metadata,
		},
		Thread: openai.ThreadRequest{
			Messages: []openai.ThreadMessage{
				{
					Role:    "user",
					Content: msg,
				},
			},
		},
	})
}

// 创建run，意思是把线程放在这个上面跑
func createRun(ctx context.Context, threadID, assistantID string) (openai.Run, error) {
	return client.CreateRun(ctx, threadID, openai.RunRequest{
		AssistantID: assistantID,
	})
}

// 查询run线程
func getRun(ctx context.Context, threadId, runId string) {
	response, err := client.RetrieveRun(ctx, threadId, runId)
	if err != nil {
		panic(err)
	}
	objToJson(response, "getRun")
}

// 取消线程
func cancelRun(ctx context.Context, threadId, runId string) {
	response, err := client.CancelRun(ctx, threadId, runId)
	if err != nil {
		panic(err)
	}
	objToJson(response, "cancelRun")
}

// 列表run线程
func listRun(ctx context.Context, threadId string) {
	response, err := client.ListRuns(ctx, threadId, openai.Pagination{})
	if err != nil {
		panic(err)
	}
	objToJson(response, "listRun")
}

// 运行步骤详情
func getRunStep(ctx context.Context, threadId, runId, stepId string) {
	response, err := client.RetrieveRunStep(ctx, threadId, runId, stepId)
	if err != nil {
		panic(err)
	}
	objToJson(response, "getRunStep")
}

// 运行步骤列表
func listRunStep(ctx context.Context, threadId, runId string) {
	response, err := client.ListRunSteps(ctx, threadId, runId, openai.Pagination{})
	if err != nil {
		panic(err)
	}
	objToJson(response, "listRunStep")
}

// 模拟ai对话
// 创建助手
// 创建线程
// 创建message放在线程中
// 创建运行线程
// 查询信息
func TestThreadRun(t *testing.T) {
	ctx := context.Background()
	assistantId := "asst_HUMjqIbCJ2DXZsqnuPaFIhwY"
	threadId := "thread_wJOlItwVmYRQlBkYhX83ZAsj"

	_, err := createMessage(ctx, threadId, "这个题是用go实现解方程“3x+11=14”")
	if err != nil {
		panic(err)
	}

	run, err := createRun(ctx, assistantId, threadId)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		getRun(ctx, threadId, run.ID)
		time.Sleep(1 * time.Second)
	}

	listMessage(ctx, threadId)

	listRun(ctx, threadId)
}
