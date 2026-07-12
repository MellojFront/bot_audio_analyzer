package telegram

import (
	"context"
	"fmt"
)

func (c *Client) GetUpdates(ctx context.Context, offset int64) ([]Update, error) {

	requestBody := GetUpdatesRequest{
		Offset:  offset,
		Timeout: 30,
	}

	var response GetUpdatesResponse
	if err := c.doJSON(ctx, "getUpdates", requestBody, &response); err != nil {
		return nil, fmt.Errorf("get Telegram updates: %w", err)
	}

	if !response.OK {
		return nil, fmt.Errorf("Telegram API error %d: %s", response.ErrorCode, response.Description)
	}

	return response.Result, nil
}
