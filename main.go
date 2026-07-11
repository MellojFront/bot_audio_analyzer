package main

import (
	"bot_audio_analyzer/app/analyzer"
	"bot_audio_analyzer/app/exporter"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("укажите путь к аудиофайлу")
	}

	path := os.Args[1]

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("файл не найден: %s", path)
		}

		log.Fatalf("не удалось открыть файл: %v", err)
	}

	info, err := analyzer.Analyze(path)
	if err != nil {
		log.Fatal(err)
	}

	jsonPath, err := exporter.ToJSON(info, path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Файл:", info.Name)
	fmt.Println("Длительность:", info.Duration)
	fmt.Println("Битрейт:", info.Bitrate)
	fmt.Println("Формат:", info.Format)
	fmt.Println("Sample Rate:", info.SampleRate)
	fmt.Println("Каналы:", info.Channels)
	fmt.Println("Кодек:", info.Codec)
	fmt.Println("Битовая глубина:", info.BitDepth)
	fmt.Println("Размер файла", info.FileSize)
	fmt.Println("Громкость:", info.LUFS)
	fmt.Println("True Peak:", info.TruePeak)
	fmt.Println("Диапазон громкости:", info.LRA)
	fmt.Println("Waveform:", info.WaveformPath)
	fmt.Println("Spectrogram:", info.SpectrogramPath)
	fmt.Println("JSON:", jsonPath)

}
