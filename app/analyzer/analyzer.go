package analyzer

import (
	"bot_audio_analyzer/app/ffmpeg"
	"bot_audio_analyzer/app/formatter"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Analyze(path string) (*TrackInfo, error) {
	probe, err := ffmpeg.Probe(path)
	if err != nil {
		return nil, err
	}

	loudness, err := ffmpeg.AnalyzeLoudness(path)
	if err != nil {
		return nil, err
	}

	file, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("get file information: %w", err)
	}

	waveformPath, err := buildWaveformPath(path)
	if err != nil {
		return nil, err
	}

	if err := ffmpeg.GenerateWaveform(path, waveformPath); err != nil {
		return nil, err
	}

	var audioStream *ffmpeg.Stream

	for i := range probe.Streams {
		if probe.Streams[i].CodecType == "audio" {
			audioStream = &probe.Streams[i]
			break
		}
	}

	if audioStream == nil {
		return nil, fmt.Errorf("audio stream not found")
	}

	info := &TrackInfo{
		Name:       probe.Format.Filename,
		Duration:   formatter.Duration(probe.Format.Duration),
		Bitrate:    formatter.Bitrate(probe.Format.BitRate),
		Format:     probe.Format.FormatName,
		Codec:      audioStream.CodecName,
		SampleRate: audioStream.SampleRate + "Hz",
		Channels:   formatter.Channels(audioStream.Channels),
		BitDepth:   formatter.BitDepth(audioStream.BitsPerSample, audioStream.BitsPerRawSample),
		FileSize:   formatter.FileSize(file.Size()),

		LUFS:         loudness.InputI + " LUFS",
		TruePeak:     loudness.InputTP + " dBTP",
		LRA:          loudness.InputLRA + " LU",
		WaveformPath: waveformPath,
	}
	return info, nil
}

func buildWaveformPath(inputPath string) (string, error) {
	const outputDir = "output"

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("create output directory: %w", err)
	}

	fileName := filepath.Base(inputPath)
	extension := filepath.Ext(fileName)
	nameWithoutExtension := strings.TrimSuffix(fileName, extension)

	outputFileName := nameWithoutExtension + "_waveform.png"

	return filepath.Join(outputDir, outputFileName), nil
}
