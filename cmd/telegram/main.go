package main

import (
	"bot_audio_analyzer/app/telegram"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	client, err := telegram.NewClient(token)

	if err != nil {
		log.Fatal(err)
	}

	offset := int64(0)

	for {
		updates, err := client.GetUpdates(context.Background(), offset)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, update := range updates {
			offset = update.UpdateID + 1

			if update.Message == nil {
				continue
			}

			message := update.Message
			chatID := message.Chat.ID

			switch {
			case message.Text == "/start":
				markdown := `# Audio Analyzer

Отправь мне аудиофайл в формате:

- WAV
- MP3
- FLAC
- M4A
- AAC
- OGG

Я отправлю:

- технические параметры;
- LUFS;
- True Peak;
- LRA;
- waveform;
- спектрограмму.`

				if err := client.SendRichMessage(
					context.Background(),
					chatID,
					markdown,
				); err != nil {
					log.Printf("send /start response: %v", err)
				}

			case message.Audio != nil || message.Document != nil:
				if err := client.SendRichMessage(
					context.Background(),
					chatID,
					"## Анализ начат\n\nФайл получен и обрабатывается.",
				); err != nil {
					log.Printf("send processing message: %v", err)
				}

				if err := client.HandleAudioMessage(
					context.Background(),
					message,
				); err != nil {
					log.Printf("handle audio: %v", err)

					if sendErr := client.SendRichMessage(
						context.Background(),
						chatID,
						"## Ошибка\n\nНе удалось обработать аудиофайл.",
					); sendErr != nil {
						log.Printf("send error response: %v", sendErr)
					}
				}

			default:
				if err := client.SendRichMessage(
					context.Background(),
					chatID,
					"Отправь аудиофайл или используй команду **/start**.",
				); err != nil {
					log.Printf("send default response: %v", err)
				}
			}
		}
	}
}
