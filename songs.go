package main

import (
	"encoding/json"
	"github.com/edemond/stereo/db"
	"github.com/edemond/stereo/log"
	"github.com/edemond/stereo/player"
	"net/http"
)

func handleSongs(w http.ResponseWriter, req *http.Request) {
	songs, err := db.GetAllSongs()
	if err != nil {
		log.Errorf("Failed to get all songs: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(songs)
	if err != nil {
		log.Errorf("Failed to marshal songs to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header()["Access-Control-Allow-Origin"] = []string{"http://localhost:3000"}
	w.Write(bytes)
}

func handleArtists(w http.ResponseWriter, req *http.Request) {
	artists, err := db.GetArtists()
	if err != nil {
		log.Errorf("Failed to get all artists: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(artists)
	if err != nil {
		log.Errorf("Failed to marshal artists to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header()["Access-Control-Allow-Origin"] = []string{"http://localhost:3000"}
	w.Write(bytes)
}

func handlePlaySong(w http.ResponseWriter, req *http.Request) {
	log.Infof("Play song!")
	id := req.URL.Query().Get("id")
	if len(id) <= 0 {
		log.Errorf("No ID provided.")
		http.Error(w, "No ID provided.", http.StatusBadRequest)
		return
	}

	song, err := db.GetSongByID(id)
	if err != nil {
		log.Errorf("Unknown song ID: '%v'\n", id)
		http.Error(w, "Unknown song ID.", http.StatusNotFound)
		return
	}

	if song == nil {
		log.Errorf("Song was nil for ID: '%v'\n", id)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	player := player.GetPlayer()
	player.Stop()

	err = db.SetNowPlayingSong(song)
	if err != nil {
		log.Errorf("Couldn't set now-playing song: %v", err)
	}

	player.Play(song)
	// TODO: We should maybe do this in an event from the player, like with stop.
	BroadcastPlay(song)

	bytes, err := json.Marshal(song)
	if err != nil {
		log.Errorf("Failed to marshal now-playing song to JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header()["Access-Control-Allow-Origin"] = []string{"http://localhost:3000"}
	w.Write(bytes)
}

func handleStop(w http.ResponseWriter, req *http.Request) {
	log.Infof("Stop song!")
	err := db.ClearNowPlayingSong()
	if err != nil {
		log.Errorf("Couldn't clear now-playing song: %v", err)
	}

	player := player.GetPlayer()
	player.Stop()

	w.Header()["Access-Control-Allow-Origin"] = []string{"http://localhost:3000"}
}
