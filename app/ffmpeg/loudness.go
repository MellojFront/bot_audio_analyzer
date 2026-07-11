package ffmpeg

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func AnalyzeLoudness(path string) (*LoudnessResponse, error) {
	cmd := exec.Command("ffmpeg", "-i", path, "-af", "loudnorm=print_format=json", "-f", "null", "-")

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("analyze loudness: %w\n%s", err, out)
	}

	jsonData, err := extractLoudnessJSON(string(out))
	if err != nil {
		return nil, err
	}

	var loudness LoudnessResponse

	if err := json.Unmarshal([]byte(jsonData), &loudness); err != nil {
		return nil, fmt.Errorf("parse loudness json: %w", err)
	}

	return &loudness, nil
}

func extractLoudnessJSON(output string) (string, error) {
	start := strings.LastIndex(output, "{")
	end := strings.LastIndex(output, "}")

	if start == -1 || end == -1 || end < start {
		return "", fmt.Errorf("loudness json not found")
	}
	return output[start : end+1], nil
}
