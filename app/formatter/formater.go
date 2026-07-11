package formatter

import (
	"fmt"
	"strconv"
	"time"
)

func Duration(value string) string {
	second, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return value
	}
	duration := time.Duration(second * float64(time.Second))

	minutes := int(duration.Minutes())
	secondsInt := int(duration.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d", minutes, secondsInt)
}

func Bitrate(value string) string {
	bitrate, err := strconv.Atoi(value)
	if err != nil {
		return value
	}

	return fmt.Sprintf("%d kbps", bitrate/1000)
}

func Channels(channels int) string {
	switch channels {
	case 1:
		return "mono"
	case 2:
		return "stereo"
	case 6:
		return "5.1 Surround"
	case 8:
		return "7.1 Surround"
	default:
		return fmt.Sprintf("%d channels", channels)
	}
}

func BitDepth(bitsPerSample int, bitsPerRawSample string) string {
	if bitsPerSample > 0 {
		return fmt.Sprintf("%d bit", bitsPerSample)
	}

	if bitsPerRawSample != "" && bitsPerRawSample != "0" {
		return bitsPerRawSample + " bit"
	}

	return "неизвестно"
}
