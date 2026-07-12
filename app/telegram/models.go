package telegram

import "encoding/json"

type InputRichMessage struct {
	Markdown            string `json:"markdown"`
	HTML                string `json:"html"`
	IsRTL               bool   `json:"is_rtl"`
	SkipEntiryDetection bool   `json:"skip_entiry_detection"`
}

type SendRichMessageRequest struct {
	ChatID      int64            `json:"chat_id"`
	RichMessage InputRichMessage `json:"rich_message"`
}

type APIResponse struct {
	OK          bool            `json:"ok"`
	Description string          `json:"description"`
	ErrorCode   int             `json:"error_code"`
	Result      json.RawMessage `json:"result"`
}

type GetUpdatesRequest struct {
	Offset  int64 `json:"offset,omitempty"`
	Timeout int   `json:"timeout,omitempty"`
}

type GetUpdatesResponse struct {
	OK          bool     `json:"ok"`
	Result      []Update `json:"result"`
	ErrorCode   int      `json:"error_code,omitempty"`
	Description string   `json:"description,omitempty"`
}

type Update struct {
	UpdateID int64    `json:"update_id"`
	Message  *Message `json:"message,omitempty"`
}

type Message struct {
	MessageID int64  `json:"message_id"`
	Text      string `json:"text,omitempty"`
	Chat      Chat   `json:"chat"`

	Audio    *Audio    `json:"audio,omitempty"`
	Document *Document `json:"document,omitempty"`
}

type Chat struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

// save content on telegram file model
type Audio struct {
	FileID    string `json:"file_id"`
	FileName  string `json:"file_name,omitempty"`
	MimeType  string `json:"mime_type,omitempty"`
	FileSize  int64  `json:"file_size,omitempty"`
	Duration  int    `json:"duration,omitempty"`
	Title     string `json:"title,omitempty"`
	Performer string `json:"performer,omitempty"`
}

type Document struct {
	FileID   string `json:"file_id"`
	FileName string `json:"file_name,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize int64  `json:"file_size,omitempty"`
}

// model getFile
type GetFileRequest struct {
	FileID string `json:"file_id"`
}

type GetFileResponse struct {
	OK          bool         `json:"ok"`
	Result      TelegramFile `json:"result"`
	ErrorCode   int          `json:"error_code,omitempty"`
	Description string       `json:"description,omitempty"`
}

type TelegramFile struct {
	FileID   string `json:"file_id"`
	FilePath string `json:"file_path"`
	FileSize int64  `json:"file_size,omitempty"`
}
