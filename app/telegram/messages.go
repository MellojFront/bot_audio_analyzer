package telegram

import (
	"fmt"
	"path/filepath"

	"bot_audio_analyzer/app/analyzer"
)

func BuildAnalysisMessage(info *analyzer.TrackInfo) string {
	return fmt.Sprintf(
		`# Анализ аудиофайла

| Параметр | Значение |
|---|---|
| Файл | %s |
| Длительность | %s |
| Размер | %s |
| Формат | %s |
| Кодек | %s |
| Битовая глубина | %s |
| Битрейт | %s |
| Sample Rate | %s |
| Каналы | %s |
| LUFS | %s |
| True Peak | %s |
| LRA | %s |`,
		filepath.Base(info.Name),
		info.Duration,
		info.FileSize,
		info.Format,
		info.Codec,
		info.BitDepth,
		info.Bitrate,
		info.SampleRate,
		info.Channels,
		info.LUFS,
		info.TruePeak,
		info.LRA,
	)
}
