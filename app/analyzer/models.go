package analyzer

type TrackInfo struct {
	Name       string
	Duration   string
	Bitrate    string
	Format     string
	Codec      string
	BitDepth   string
	FileSize   string
	SampleRate string
	Channels   string

	LUFS     string
	TruePeak string
	LRA      string
}
