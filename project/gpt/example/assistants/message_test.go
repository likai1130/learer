package assistants

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"testing"
)

// 创建消息 TODO FileIds Metadata
func createMessage(ctx context.Context, threadId, msg string) (openai.Message, error) {
	return client.CreateMessage(ctx, threadId, openai.MessageRequest{
		Role:    "user",
		Content: msg,
		//FileIds:  nil,
		//Metadata: nil,
	})
}

// 查询消息
func getMessage(ctx context.Context, threadId, messageId string) {
	response, err := client.RetrieveMessage(ctx, threadId, messageId)
	if err != nil {
		panic(err)
	}
	objToJson(response, "getMessage")
}

// 消息列表
func listMessage(ctx context.Context, threadId string) {
	response, err := client.ListMessage(ctx, threadId, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	objToJson(response, "listMessage")
}

func TestCRUDMessage(t *testing.T) {
	ctx := context.Background()
	// 创建线程
	/*thread, err := createThread(ctx)
	if err != nil {
		panic(err)
	}*/
	threadId := "thread_KuFgsVfhPwCOpLK4Demmt5yI"
	// 创建消息

	msg, err := createMessage(ctx, threadId, "我上个问题是什么？")
	if err != nil {
		panic(err)
	}

	msgId := msg.ID

	// 查询消息
	getMessage(ctx, threadId, msgId) //{"id":"msg_bXi5oCF57DD7kydHqQ1LrbZD","object":"thread.message","created_at":1701342008,"thread_id":"thread_KuFgsVfhPwCOpLK4Demmt5yI","role":"user","content":[{"type":"text","text":{"value":"你是我的数学老师吗","annotations":[]}}],"file_ids":[],"metadata":{}}

	// 消息列表
	listMessage(ctx, threadId)
}
