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

			chatID := update.Message.Chat.ID
			text := update.Message.Text

			switch text {
			case "/start":
				markdown := `# Audio Analyzer

Отправь мне аудиофайл, и я выполню анализ:

- технические параметры;
- LUFS;
- True Peak;
- LRA;
- waveform;
- спектрограмма.`

				if err := client.SendRichMessage(
					context.Background(),
					chatID,
					markdown,
				); err != nil {
					log.Printf("send /start response: %v", err)
				}

			default:
				markdown := `Я пока понимаю только команду **/start**.

Следующий этап — приём аудиофайлов.`

				if err := client.SendRichMessage(
					context.Background(),
					chatID,
					markdown,
				); err != nil {
					log.Printf("send default response: %v", err)
				}
			}
		}
	}
	/*
				chatIDValue := os.Getenv("TELEGRAM_CHAT_ID")
				if chatIDValue == "" {
					log.Fatal("TELEGRAM_CHAT_ID is not set")
				}

				chatID, err := strconv.ParseInt(chatIDValue, 10, 64)
				if err != nil {
					log.Fatalf("invalid TELEGRAM_CHAT_ID: %v", err)
				}

				client, err := telegram.NewClient(token)
				if err != nil {
					log.Fatal(err)
				}

			markdown := `# Анализ аудиофайла

		| Параметр | Значение |
		|:---|---:|
		| Формат | WAV |
		| Кодек | PCM S16LE |
		| Битовая глубина | 16 bit |
		| Частота | 44.1 kHz |
		| Каналы | Stereo |
		| LUFS | -9.42 |
		| True Peak | -0.31 dBTP |
		| LRA | 4.20 LU |`
	*/

}
