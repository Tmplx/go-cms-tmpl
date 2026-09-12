package paddle

import (
	paddle "github.com/PaddleHQ/paddle-go-sdk/v5"
)

type PaddleClient struct {
	Client *paddle.SDK
}

func NewPaddleClient(apiKey string) (*PaddleClient, error) {
	client, err := paddle.New(apiKey, paddle.WithBaseURL(paddle.SandboxBaseURL))
	if err != nil {
		return nil, err
	}
	return &PaddleClient{
		Client: client,
	}, nil
}
