package ffmpeg

import (
	"fmt"
	"os/exec"
)

func GenerateWaveform(inputPath, outputPath string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", inputPath,
		"-filter_complex", "showwavespic=s=1200x300",
		"-frames:v", "1",
		outputPath,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("generate waveform: %w\n%s", err, out)
	}

	return nil
}
