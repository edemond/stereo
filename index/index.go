package index

import(
  "fmt"
  "github.com/edemond/stereo/log"
  "github.com/edemond/stereo/media"
  "os"
  "path"
  "strings"
)

var _audioFileTypes = map[string]bool {
  ".flac": true,
  ".m3u": true,
  ".m4a": true,
  ".mp3": true,
  ".ogg": true, 
  ".wav": true,
  ".wma": true,
}

// SongDetailer inspects a song file and returns details like artist, title, and length.
type SongDetailer interface {
  GetSongDetails(file os.DirEntry) (*media.SongFile, error)
  GetArtistDetails(file os.DirEntry) (*media.Artist, error)
}

// Index indexes a directory, returning a list of media.
func Index(path string) (*media.DirectoryContents, error) {
  files, err := getFilesRecursive(path)
  if err != nil {
    return nil, fmt.Errorf("Couldn't index %v: %v", path, err)
  }

  // TODO: Support a batting order of different detail-reading strategies.
  // Different file types have different kinds of metadata, so file type plays into it.
  details := &SongDetailsFromFilename{}

  allArtistSongs := map[string]*media.ArtistSongs{}

  for _,file := range files {
    song, err := details.GetSongDetails(file, path)
    if err != nil {
      log.Warningf("Couldn't create a song from file %v: %v", file.Name(), err)
      continue
    }
    
    artistSongs, ok := allArtistSongs[song.Artist]
    if !ok {
      allArtistSongs[song.Artist] = &media.ArtistSongs{
        Artist: &media.Artist{
          Name: song.Artist,
        },
        Songs: []*media.SongFile{},
      }
    } else {
      artistSongs.Songs = append(artistSongs.Songs, song)
    }
  }

  artistSongsList := []*media.ArtistSongs{}
  for _,a := range allArtistSongs {
    artistSongsList = append(artistSongsList, a)
  }

  return &media.DirectoryContents{
    ArtistSongs: artistSongsList,
  }, nil
}

func getFilesRecursive(path string) ([]os.DirEntry, error) {
  entries, err := os.ReadDir(path)
  if err != nil {
    return nil, fmt.Errorf("Couldn't list files in %v: %v", path, err)
  }

  files := []os.DirEntry{}

  for _,entry := range entries {
    // TODO: maybe don't follow symlinks
    if entry.IsDir() {
      childFiles, err := getFilesRecursive(entry.Name())
      if err != nil {
        log.Warningf("Failed to index %v: %v", path, err)
        continue
      } 
      files = append(files, childFiles...)
    }
    if shouldIndexFile(entry) {
      log.Infof("Indexing: %v", entry.Name())
      files = append(files, entry)
    }
  }

  return files, nil
}

func shouldIndexFile(file os.DirEntry) bool {
  return isAudioFile(file.Name())
}

func isAudioFile(filename string) bool {
  extension := strings.ToLower(path.Ext(filename))
  _,ok := _audioFileTypes[extension]
  return ok
}
