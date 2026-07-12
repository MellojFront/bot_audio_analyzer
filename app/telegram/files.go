package telegram

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (c *Client) GetFile(
	ctx context.Context,
	fileID string,
) (*TelegramFile, error) {
	requestBody := GetFileRequest{
		FileID: fileID,
	}

	var response GetFileResponse

	if err := c.doJSON(
		ctx,
		"getFile",
		requestBody,
		&response,
	); err != nil {
		return nil, fmt.Errorf("get Telegram file: %w", err)
	}

	if !response.OK {
		return nil, fmt.Errorf(
			"Telegram API error %d: %s",
			response.ErrorCode,
			response.Description,
		)
	}

	return &response.Result, nil
}

func (c *Client) DownloadFile(
	ctx context.Context,
	filePath string,
	outputPath string,
) error {
	url := fmt.Sprintf(
		"https://api.telegram.org/file/bot%s/%s",
		c.token,
		filePath,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return fmt.Errorf("create file download request: %w", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("download Telegram file: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"download Telegram file: status %s",
			response.Status,
		)
	}

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("create download directory: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create downloaded file: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, response.Body); err != nil {
		return fmt.Errorf("save downloaded file: %w", err)
	}

	return nil
}
