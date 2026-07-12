package telegram

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"bot_audio_analyzer/app/analyzer"
)

func (c *Client) HandleAudioMessage(
	ctx context.Context,
	message *Message,
) error {
	fileID, fileName, err := extractAudioFile(message)
	if err != nil {
		return err
	}

	telegramFile, err := c.GetFile(ctx, fileID)
	if err != nil {
		return err
	}

	tempDir, err := os.MkdirTemp("", "audio-analyzer-*")
	if err != nil {
		return fmt.Errorf("create temporary directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	if fileName == "" {
		fileName = filepath.Base(telegramFile.FilePath)
	}

	fileName = filepath.Base(fileName)

	localPath := filepath.Join(tempDir, fileName)

	if err := c.DownloadFile(
		ctx,
		telegramFile.FilePath,
		localPath,
	); err != nil {
		return err
	}

	info, err := analyzer.Analyze(localPath)
	if err != nil {
		return fmt.Errorf("analyze audio: %w", err)
	}

	report := BuildAnalysisMessage(info)

	if err := c.SendRichMessage(
		ctx,
		message.Chat.ID,
		report,
	); err != nil {
		return err
	}

	if err := c.SendPhoto(
		ctx,
		message.Chat.ID,
		info.WaveformPath,
		"Waveform",
	); err != nil {
		return err
	}

	if err := c.SendPhoto(
		ctx,
		message.Chat.ID,
		info.SpectrogramPath,
		"Спектрограмма",
	); err != nil {
		return err
	}

	return nil
}

func extractAudioFile(message *Message) (
	fileID string,
	fileName string,
	err error,
) {
	if message.Audio != nil {
		return message.Audio.FileID, message.Audio.FileName, nil
	}

	if message.Document != nil {
		if !isSupportedAudioDocument(message.Document) {
			return "", "", fmt.Errorf(
				"unsupported document type: %s",
				message.Document.MimeType,
			)
		}

		return message.Document.FileID, message.Document.FileName, nil
	}

	return "", "", fmt.Errorf("audio file not found in message")
}

func isSupportedAudioDocument(document *Document) bool {
	if strings.HasPrefix(document.MimeType, "audio/") {
		return true
	}

	extension := strings.ToLower(
		filepath.Ext(document.FileName),
	)

	switch extension {
	case ".wav", ".mp3", ".flac", ".m4a", ".aac", ".ogg":
		return true
	default:
		return false
	}
}
