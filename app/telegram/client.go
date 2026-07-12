package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const defaultBaseURL = "https://api.telegram.org"

type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

func NewClient(token string) (*Client, error) {
	token = strings.TrimSpace(token)
	if token == "" {
		return nil, fmt.Errorf("telegram bot token is empty")
	}

	return &Client{
		token:   token,
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: 40 * time.Second,
		},
	}, nil
}

func (c *Client) SendRichMessage(ctx context.Context, chatID int64, markdown string) error {
	requestBody := SendRichMessageRequest{
		ChatID: chatID,
		RichMessage: InputRichMessage{
			Markdown: markdown,
		},
	}

	var response APIResponse

	if err := c.doJSON(ctx, "sendRichMessage", requestBody, &response); err != nil {
		return err
	}

	if !response.OK {
		return fmt.Errorf("telegram API error %d: %s", response.ErrorCode, response.Description)
	}
	return nil
}

func (c *Client) doJSON(ctx context.Context, method string, requestBody any, responseBody any) error {
	data, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("encode telegram request: %w", err)
	}

	url := fmt.Sprintf(
		"%s/bot%s/%s",
		c.baseURL,
		c.token,
		method,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewReader(data),
	)

	if err != nil {
		return fmt.Errorf("create telegram request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("send telegram request: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("read telegram response: %w", err)
	}

	if err := json.Unmarshal(body, responseBody); err != nil {
		return fmt.Errorf(
			"decode telegram response: %w; body: %s",
			err,
			body,
		)
	}

	return nil
}
