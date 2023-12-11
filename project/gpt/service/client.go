package service

import (
	"github.com/sashabaranov/go-openai"
	"sync"
)

var client *openai.Client
var once sync.Once

func init() {
	createOpenAIClient("")
}
func createOpenAIClient(token string) *openai.Client {
	once.Do(func() {
		client = openai.NewClient(token)
	})
	return client
}
