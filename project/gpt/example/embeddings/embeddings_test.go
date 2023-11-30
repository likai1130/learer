package embeddings

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"testing"
)

func TestEmbeddings(t *testing.T) {
	c := openai.NewClient("")
	ctx := context.Background()
	res, err := c.CreateEmbeddings(ctx, openai.EmbeddingRequest{
		Input: []string{
			"Israel Gaza: Three-year-old-twins among hostages released by Hamas",
		},
		Model:          openai.AdaEmbeddingV2,
		EncodingFormat: openai.EmbeddingEncodingFormatBase64,
	})
	if err != nil {
		panic(err)
	}

	marshal, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))
}
