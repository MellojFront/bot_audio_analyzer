package ffmpeg

import (
	"fmt"
	"os/exec"
)

func GenerateSpectrogram(inputPath, outputPath string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", inputPath,
		"-lavfi", "showspectrumpic=s=1200x600:legend=disabled",
		outputPath,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("generate spectrogram: %w\n%s", err, out)
	}
	return nil
}
