package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://api.jsonbin.io/v3/b"

type Client struct {
	Key string
}

func NewClient(apiKey string) *Client {
	return &Client{Key: apiKey}
}

type createResponse struct {
	Metadata struct {
		ID string `json:"id"`
	} `json:"metadata"`
}

func (c *Client) Create(data []byte, name string) (string, error) {
	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", c.Key)
	req.Header.Set("X-Bin-Name", name)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var result createResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	return result.Metadata.ID, nil
}

func (c *Client) Get(id string) (string, error) {
	req, err := http.NewRequest("GET", baseURL+"/"+id, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Master-Key", c.Key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}
	return string(body), nil
}

func (c *Client) Update(id string, data []byte) error {
	req, err := http.NewRequest("PUT", baseURL+"/"+id, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", c.Key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}
	return nil
}

func (c *Client) Delete(id string) error {
	req, err := http.NewRequest("DELETE", baseURL+"/"+id, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Master-Key", c.Key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}
	return nil
}
