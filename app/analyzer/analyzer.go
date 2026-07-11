package analyzer

import (
	"bot_audio_analyzer/app/ffmpeg"
	"bot_audio_analyzer/app/formatter"
	"fmt"
)

func Analyze(path string) (*TrackInfo, error) {
	probe, err := ffmpeg.Probe(path)
	if err != nil {
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
		BitDepth: formatter.BitDepth(
			audioStream.BitsPerSample,
			audioStream.BitsPerRawSample,
		),
	}
	return info, nil
}
