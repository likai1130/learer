package assistants

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"testing"
)

/**
测试助手的使用
Assistants	API 旨在帮助开发人员构建能够执行各种任务的强大 AI 助手。

Assistant	使用 OpenAI模型和调用工具的专用 AI
Thread		助理和用户之间的对话会话。线程存储消息并自动处理截断以使内容适合模型的上下文。
Message		由助理或用户创建的消息。消息可以包括文本、图像和其他文件。消息以列表形式存储在线程上。
Run			在线程上调用助手。助手使用它的配置和线程的消息通过调用模型和工具来执行任务。作为运行的一部分，助手将消息附加到线程。
RunStep		助理在运行过程中所采取的步骤的详细列表。助手可以在运行期间调用工具或创建消息。检查运行步骤可以让您反思助手如何获得最终结果。
*/

var client *openai.Client

func init() {
	client = openai.NewClient("")

}

/*
*
创建助手
*/
func CreateAssistants(ctx context.Context) (openai.Assistant, error) {
	name := "数学导师"
	instructions := "您是一名私人数学导师。编写并运行GO代码来回答数学问题。"
	description := "数学导师无所不知"

	// 1. 创建助手
	return client.CreateAssistant(ctx, openai.AssistantRequest{
		Model:        openai.GPT3Dot5Turbo1106,
		Name:         &name,
		Description:  &description,
		Instructions: &instructions,
	})
}

/*
*
查询助手

params assistantsId 助手id
*/
func GetAssistants(ctx context.Context, assistantsId string) {
	response, err := client.RetrieveAssistant(ctx, assistantsId)
	if err != nil {
		panic(err)
	}
	objToJson(response, "GetAssistants")
}

/*
*
列表
*/
func ListAssistants(ctx context.Context) {
	response, err := client.ListAssistants(ctx, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	objToJson(response, "ListAssistants")
}

func DeleteAssistants(ctx context.Context, assistantsId string) {
	response, err := client.DeleteAssistant(ctx, assistantsId)
	if err != nil {
		panic(err)
	}
	objToJson(response, "DeleteAssistants")
}

func objToJson(obj interface{}, funcName string) {
	marshal, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("%s result = %s", funcName, string(marshal)))
}

func TestCRUDAssistants(t *testing.T) {
	ctx := context.Background()
	// 创建
	assistants, err := CreateAssistants(ctx)
	if err != nil {
		panic(err)
	}
	objToJson(assistants, "CreateAssistants")
	// 查询
	assistantId := assistants.ID
	GetAssistants(ctx, assistantId)
	//列表
	ListAssistants(ctx)
	//删除
	DeleteAssistants(ctx, assistantId)
}

func TestAssistantsList(t *testing.T) {
	ctx := context.Background()
	ListAssistants(ctx)
}

func TestAssistants(t *testing.T) {
	ctx := context.Background()
	name := "数学导师"
	instructions := "您是一名私人数学导师。编写并运行GO代码来回答数学问题。"
	description := "数学导师无所不知"

	// 1. 创建助手
	response, err := client.CreateAssistant(ctx, openai.AssistantRequest{
		Model:        openai.GPT3Dot5Turbo1106,
		Name:         &name,
		Description:  &description,
		Instructions: &instructions,
	})

	if err != nil {
		panic(err)
	}

	marshal, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))

	// 创建线程
	resp, err := client.CreateThread(ctx, openai.ThreadRequest{
		Messages: []openai.ThreadMessage{
			{
				Role:    openai.ThreadMessageRoleUser,
				Content: "我需要解方程“3x+11=14”。你能帮我吗？",
			},
		},
	})

	if err != nil {
		panic(err)
	}
	marshal, err = json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))

	// 3. 向主题添加信息
	msgResp, err := client.CreateMessage(ctx, resp.ID, openai.MessageRequest{
		Role:    "user",
		Content: "我需要解方程“3x+11=14”。你能帮我吗?",
	})
	if err != nil {
		panic(err)
	}

	marshal, err = json.Marshal(msgResp)
	if err != nil {
		panic(err)
	}
	fmt.Println(marshal)
	// 4. 运行助手
	run, err := client.CreateThreadAndRun(ctx, openai.CreateThreadAndRunRequest{
		RunRequest: openai.RunRequest{
			AssistantID:  response.ID,
			Model:        &response.Model,
			Instructions: &instructions,
		},
		Thread: openai.ThreadRequest{
			Messages: []openai.ThreadMessage{
				{
					Role:    "user",
					Content: "我需要解方程“3x+11=14”。你能帮我吗?",
				},
			},
			Metadata: nil,
		},
	})

	if err != nil {
		panic(err)
	}

	marshal, err = json.Marshal(run)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))

	// 5. 检查运行状态
	// 6. 显示助力的响应
	messages, err := client.ListMessage(ctx, resp.ID, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	marshal, err = json.Marshal(messages)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}
