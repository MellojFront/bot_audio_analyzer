package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func (c *Client) SendPhoto(
	ctx context.Context,
	chatID int64,
	imagePath string,
	caption string,
) error {
	file, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("open photo: %w", err)
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	if err := writer.WriteField(
		"chat_id",
		fmt.Sprintf("%d", chatID),
	); err != nil {
		return fmt.Errorf("write chat_id field: %w", err)
	}

	if caption != "" {
		if err := writer.WriteField("caption", caption); err != nil {
			return fmt.Errorf("write caption field: %w", err)
		}
	}

	part, err := writer.CreateFormFile(
		"photo",
		filepath.Base(imagePath),
	)
	if err != nil {
		return fmt.Errorf("create photo form field: %w", err)
	}

	if _, err := io.Copy(part, file); err != nil {
		return fmt.Errorf("write photo data: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("close multipart writer: %w", err)
	}

	url := fmt.Sprintf(
		"%s/bot%s/sendPhoto",
		c.baseURL,
		c.token,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		&body,
	)
	if err != nil {
		return fmt.Errorf("create sendPhoto request: %w", err)
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("send photo: %w", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("read sendPhoto response: %w", err)
	}

	var apiResponse APIResponse

	if err := json.Unmarshal(responseBody, &apiResponse); err != nil {
		return fmt.Errorf(
			"decode sendPhoto response: %w; body: %s",
			err,
			responseBody,
		)
	}

	if !apiResponse.OK {
		return fmt.Errorf(
			"Telegram API error %d: %s",
			apiResponse.ErrorCode,
			apiResponse.Description,
		)
	}

	return nil
}
