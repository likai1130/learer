package embeddings

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"sync"
	"testing"
)

type VectorRequest struct {
	Text string `json:"text"`
}

type VectorResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Vector struct {
	Text string
	c    *openai.Client
	resp chan VectorResponse
}

func NewVector(text string) *Vector {
	c := openai.NewClient("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")
	return &Vector{
		Text: text,
		c:    c,
		resp: make(chan VectorResponse),
	}
}

func (v *Vector) Embeddings() VectorResponse {
	ctx := context.Background()
	res, err := v.c.CreateEmbeddings(ctx, openai.EmbeddingRequest{
		Input: []string{
			v.Text,
		},
		Model:          openai.AdaEmbeddingV2,
		EncodingFormat: openai.EmbeddingEncodingFormatBase64,
	})
	if err != nil {
		return VectorResponse{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		}
	}
	return VectorResponse{
		Code: http.StatusOK,
		Data: res,
	}
}

func processRequest(q <-chan *Vector) {
	for req := range q {
		embeddings := req.Embeddings()
		req.resp <- embeddings
		close(req.resp)
	}
}

var vectorQueue = make(chan *Vector, 5)

// 不支持并发，用队列
func TestQueueEmbeddings(t *testing.T) {
	go processRequest(vectorQueue)

	text := "Israel Gaza: Three-year-old-twins among hostages released by Hamas"
	// 模拟10个请求，用队列

	var wg sync.WaitGroup
	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			content := fmt.Sprintf("%s %d", text, i)
			vector := NewVector(content)
			vectorQueue <- vector
			resp := <-vector.resp
			marshal, _ := json.Marshal(resp)
			fmt.Println(string(marshal))
		}(i, &wg)
	}
	wg.Wait()
	fmt.Println("All requests processed")
}
