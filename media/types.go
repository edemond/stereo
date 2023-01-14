package media

import(
  "fmt"
)

// The result of indexing a directory.
type DirectoryContents struct {
  ArtistSongs []*ArtistSongs
}

type ArtistSongs struct {
  Artist *Artist
  Songs []*SongFile
}

type Song struct {
  Id string `json:"id"`
  ArtistId string `json:"artistId"`
  Title string `json:"title"`
  Artist string `json:"artist"`
}

type Artist struct {
  Id string `json:"id"`
  Name string `json:"name"`
}

// Server-side representation of the song for playback purposes.
// We don't send this to the client.
type SongFile struct {
  Song
  Path string `json:"path"`
  Type string `json:"type"`
}

type SearchResult struct {
  Text string `json:"text"`
  Id string `json:"id"`
}

func (s *SongFile) GetName() string {
  return fmt.Sprintf("%v - %v", s.Artist, s.Title)
}

func (s *SongFile) GetPath() string {
  return s.Path
}
