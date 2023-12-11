package example

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
	"testing"
)

func TestProxy(t *testing.T) {
	config := openai.DefaultConfig("")
	proxyUrl, err := url.Parse("https://localhost:8820")
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}

	c := openai.NewClientWithConfig(config)
	ctx := context.Background()
	reponse, err := c.ListAssistants(ctx, nil, nil, nil, nil)

	//response, err := c.RetrieveAssistant(ctx, "asst_HUMjqIbCJ2DXZsqnuPaFIhwY")
	if err != nil {
		panic(err)
	}
	fmt.Println(reponse.Assistants)

}
