package main

import(
  "encoding/json"
  "github.com/edemond/stereo/db"
  "github.com/edemond/stereo/log"
  "net/http"
)

func handleNowPlaying(rw http.ResponseWriter, req *http.Request) {
  song, err := db.GetNowPlayingSong()
  if err != nil {
    log.Errorf("Failed to get now-playing song: %v", err)
    rw.WriteHeader(http.StatusInternalServerError)
    return
  }

  bytes, err := json.Marshal(song)
  if err != nil {
    log.Errorf("Failed to marshal now-playing song to JSON: %v", err)
    rw.WriteHeader(http.StatusInternalServerError)
    return
  }

  // TODO: Add a debug mode
  rw.Header()["Access-Control-Allow-Origin"] = []string{"http://localhost:3000"}

  rw.Write(bytes)
}

