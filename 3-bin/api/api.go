package api

type Client struct {
	Key string
}

func NewClient(apiKey string) *Client {
	return &Client{
		Key: apiKey,
	}
}
