package exporter

import (
	"bot_audio_analyzer/app/analyzer"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ToJSON(info *analyzer.TrackInfo, inputPath string) (string, error) {
	const outputDir = "output"

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("create output directory: %w", err)
	}

	fileName := filepath.Base(inputPath)
	extension := filepath.Ext(fileName)
	nameWithoutExtension := strings.TrimSuffix(fileName, extension)

	outputPath := filepath.Join(
		outputDir,
		nameWithoutExtension+"_analysis.json",
	)

	data, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return "", fmt.Errorf("encode analysis result to json: %w", err)
	}

	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		return "", fmt.Errorf("write json file: %w", err)
	}

	return outputPath, nil
}
