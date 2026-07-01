package ffmpeg

type Rendition struct {
	Name         string
	Width        int
	Height       int
	VideoBitrate string
	AudioBitrate string
	Bandwidth    int
}

var Renditions = []Rendition{
	{
		Name:         "1080p",
		Width:        1920,
		Height:       1080,
		VideoBitrate: "5000k",
		AudioBitrate: "192k",
		Bandwidth:    5000000,
	},
	{
		Name:         "720p",
		Width:        1280,
		Height:       720,
		VideoBitrate: "2800k",
		AudioBitrate: "128k",
		Bandwidth:    2800000,
	},
	{
		Name:         "480p",
		Width:        854,
		Height:       480,
		VideoBitrate: "1400k",
		AudioBitrate: "96k",
		Bandwidth:    1400000,
	},
}
