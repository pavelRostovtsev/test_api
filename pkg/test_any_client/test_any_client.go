package test_any_client

type Input struct {
	AnyData string
}

type Client struct {
	input Input
}

func NewClient(input Input) *Client {
	return &Client{input: input}
}
