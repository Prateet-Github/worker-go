package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type CompleteVideoRequest struct {
	HLSURL       string `json:"hlsUrl"`
	ThumbnailKey string `json:"thumbnailKey"`
}

func (c *Client) CompleteVideo(
	ctx context.Context,
	videoID string,
	req CompleteVideoRequest,
) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/internal/videos/%s/complete", c.baseURL, videoID),
		bytes.NewBuffer(body),
	)

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	return nil
}
