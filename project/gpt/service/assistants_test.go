package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"learner/project/gpt/service/assistants"
	"log"
	"testing"
	"time"
)

var assistantId = "asst_HUMjqIbCJ2DXZsqnuPaFIhwY" //数学导师助手
var threadId = "thread_Wk60Uo1CaJe5IoN47qy0d4jm"  //数学导师线程
var cli = createOpenAIClient("")

func TestAssistants(t *testing.T) {
	ctx := context.Background()
	// 用线程发送信息
	thread := assistants.NewThread(cli)
	/*run, err := thread.CreateAndRun(ctx, assistantId, "再帮我解2x+11=21这个方程式")
	if err != nil {
		log.Println("创建信息出错 err=", err)
		return
	}
	getThread, err := thread.GetThread(ctx, run.ThreadID)
	if err != nil {
		log.Println("查看线程出错 err=", err)
		return
	}
	bytes, err := json.Marshal(getThread)
	if err != nil {
		log.Println("getThread json err", err)
		return
	}
	log.Printf("thread 详情 = %v \n", string(bytes))*/
	newMessage := assistants.NewMessage(cli)
	_, err := newMessage.CreateMessage(ctx, threadId, "计算4x + 5 = 20这个方程")
	if err != nil {
		log.Println("createMessage  err", err)
		return
	}
	run, err := thread.Run(ctx, threadId, assistantId)
	if err != nil {
		log.Println("Run  err", err)
		return
	}

	for {
		getRun, err := thread.GetRun(ctx, run.ThreadID, run.ID)
		if err != nil {
			log.Println("for getRun线程出错 err=", err)
			return
		}
		log.Println("当前执行状态为", getRun.Status)
		if getRun.Status == openai.RunStatusCompleted {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	// 查信息列表
	message := assistants.NewMessage(cli)
	messages, err := message.ListMessages(ctx, run.ThreadID, assistants.PageTools{
		Limit: 100,
		Order: "desc",
	})
	if err != nil {
		log.Println("查询list message错误了，err = ", err)
		return
	}
	marshal, err := json.Marshal(messages)
	if err != nil {
		log.Println("json err = ", err)
		return
	}
	log.Println("消息列表：", string(marshal))

}

func TestName(t *testing.T) {
	ctx := context.Background()
	thread, err := assistants.NewThread(cli).GetThread(ctx, "thread_Wk60Uo1CaJe5IoN47qy0d4jm")
	if err != nil {
		panic(err)
	}
	marshal, err := json.Marshal(thread)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}
