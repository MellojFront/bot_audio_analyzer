package ffmpeg

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func Probe(path string) (*FFProbeResponse, error) {
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", "audio/melodic_techno_test.wav")

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("run ffprobe: %w\n%s", err, out)
	}

	var probe FFProbeResponse
	if err := json.Unmarshal(out, &probe); err != nil {
		return nil, fmt.Errorf("parse json: %w", err)
	}
	return &probe, nil
}
