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
