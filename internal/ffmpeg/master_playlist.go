package ffmpeg

import (
	"os"
	"path/filepath"
)

func (s *Service) GenerateMasterPlaylist(outputDir string) error {

	content := `#EXTM3U
#EXT-X-VERSION:3

#EXT-X-STREAM-INF:BANDWIDTH=5000000,RESOLUTION=1920x1080
1080p/index.m3u8

#EXT-X-STREAM-INF:BANDWIDTH=2800000,RESOLUTION=1280x720
720p/index.m3u8

#EXT-X-STREAM-INF:BANDWIDTH=1400000,RESOLUTION=854x480
480p/index.m3u8
`

	path := filepath.Join(outputDir, "master.m3u8")

	return os.WriteFile(path, []byte(content), 0644)
}
