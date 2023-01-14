package index

import (
	"fmt"
	"github.com/edemond/stereo/media"
	"os"
	"path"
	"regexp"
)

var _songRegexp *regexp.Regexp

func init() {
	_songRegexp = regexp.MustCompile("^(?P<artist>.+?) - (?P<title>.+?)\\.(?P<type>[a-zA-Z0-9]+)$")
}

// SongDetailsFromFilename indexes songs by guessing the artist, title, and type from the filename.
type SongDetailsFromFilename struct{}

func (d *SongDetailsFromFilename) GetSongDetails(file os.DirEntry, dir string) (*media.SongFile, error) {
	matches := _songRegexp.FindStringSubmatch(file.Name())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("Couldn't parse song details from filename '%v'", file.Name())
	}

	return &media.SongFile{
		Song: media.Song{
			Artist: matches[1],
			Title:  matches[2],
		},
		Path: path.Join(dir, file.Name()),
		Type: matches[3],
	}, nil
}
