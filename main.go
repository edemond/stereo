package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/edemond/stereo/db"
	"github.com/edemond/stereo/index"
	"github.com/edemond/stereo/log"
	"github.com/edemond/stereo/player"
	"io/fs"
	"net/http"
)

var _mediaDir string
var _port uint
var _reindex bool

// Embed the production build of the Mint application.
//go:embed ui/dist
var _static embed.FS

func init() {
	flag.StringVar(
		&_mediaDir,
		"dir",
		"/media/music",
		"Directory where media is stored.",
	)

	flag.UintVar(
		&_port,
		"port",
		8080,
		"Port for the server to listen on (default: 8080).",
	)

	flag.BoolVar(
		&_reindex,
		"reindex",
		true,
		"Whether or not to reindex songs on startup (default: true).",
	)
}

func main() {
	flag.Parse()

	log.Infof("Initializing database...")
	err := db.Initialize()
	if err != nil {
		log.Fatalf("Couldn't initialize database: %v", err)
	}

	if _reindex {
		reindex()
	}

	// TODO: Doesn't feel like the right place to hook this up.
	player := player.GetPlayer()
	player.OnSongEnd(func() {
		BroadcastStop()
	})

	// API routes (called from Mint)
	http.HandleFunc("/api/nowplaying", handleNowPlaying)
	http.HandleFunc("/api/artists", handleArtists)
	http.HandleFunc("/api/play", handlePlaySong)
	http.HandleFunc("/api/songs", handleSongs)
	http.HandleFunc("/api/stop", handleStop)
	http.HandleFunc("/api/ws", handleWebsocket)

	// Static routes (CSS, JS, images)
	fileServer := http.FileServer(getStaticFilesystem())
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	// App routes
	http.Handle("/", fileServer)

	log.Infof("Listening on localhost:%v.", _port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", _port), nil))
}

// getStaticFilesystem gets an http package filesystem that points into the
// embedded Mint files, tacking on the appropriate path.
func getStaticFilesystem() http.FileSystem {
	subFS, err := fs.Sub(_static, "ui/dist")
	if err != nil {
		log.Fatalf("Couldn't set up embedded FS for static files: %v", err)
	}
	return http.FS(subFS)
}

func reindex() {
	log.Infof("Scanning %v...", _mediaDir)
	contents, err := index.Index(_mediaDir)
	if err != nil {
		log.Infof("Failed to scan %v: %v", _mediaDir, err)
		return
	}

	log.Infof(
		"Scanning complete. Found %v artists.",
		len(contents.ArtistSongs),
	)

	log.Infof("Indexing directory contents...")
	err = db.Reindex(contents)
	if err != nil {
		log.Warningf("Failed to save directory contents to index: %v", err)
	}
}
