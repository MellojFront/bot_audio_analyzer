package analyzer

type TrackInfo struct {
	Name       string `json:"name"`
	Duration   string `json:"duration"`
	Bitrate    string `json:"bitrate"`
	Format     string `json:"format"`
	Codec      string `json:"codec"`
	BitDepth   string `json:"bit_depth"`
	FileSize   string `json:"file_size"`
	SampleRate string `json:"sample_rate"`
	Channels   string `json:"channels"`

	LUFS         string `json:"lufs"`
	TruePeak     string `json:"true_peak"`
	LRA          string `json:"lra"`
	WaveformPath string `json:"waveform_path"`

	SpectrogramPath string `json:"spectrogram_path"`
}
