package open_ai01

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"github.com/walkjohnp/go-demo/common"
	"net/http"
	"net/url"
)

var (
	token = common.GetEnv("gpt-token")

	proxyAddress = "http://localhost:7890"
)

func TestGpt() {
	client := getProxyClient()
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func getProxyClient() *openai.Client {
	config := openai.DefaultConfig(token)
	proxyUrl, err := url.Parse(proxyAddress)
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
	return c
}
