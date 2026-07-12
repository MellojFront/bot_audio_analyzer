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
	Text      string `json:"text"`
	Chat      Chat   `json:"chat"`
}

type Chat struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}
