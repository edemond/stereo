# stereo

A web frontend for a music server for my living room, intended for Raspberry Pi. Bookmark it on your phone; play songs.

It's half-finished, full of TODOs, and mostly for fun. It's also to see how it goes bolting Go and [Mint](https://mint-lang.com) together into a single-binary, single-page application.
NPM-based tooling can't be the final evolution of building these kinds of web applications, and both Go and Mint have nice, clean, simple toolchains with Swiss-army-knife commands for building and running.
It's also a continuation of an earlier, server-side version of the application (which had the nice property of being, you know, _working_ enough to listen to music on.)

## Architecture

The app has an SPA frontend in [Mint](https://mint-lang.com/) served by a backend in Go, which manages music indexing and song playback. The whole thing compiles down to a single binary, static content and all, thanks to [`go:embed`](https://pkg.go.dev/embed).

You point the Go server at a directory of music files, and it indexes them in SQLite.
It publishes API endpoints for the transport controls (i.e. for playing and stopping songs), and currently plays music in the somewhat cheesy manner of shelling out to an [MPlayer](https://en.wikipedia.org/wiki/MPlayer) process.

To add music to it, the intent is you open a Samba share on your computer and drag files over, because it's 2023 and I listen to directories of music files. Eventually the server will be watching the directory via inotify and reindex your tunes.

Multiple clients can connect at the same time and see what's playing, with the current transport state synced over websockets.

It can be served on a local network at an address like `stereo.local` using mDNS (e.g. with Avahi).

## Future ideas

- Right now we don't index or pull any metadata about the files aside from what we can glean from an overly-specific filename format.
- The Go server should eventually bake in a real music player library via cgo, if I can find a good one that can cross-compile to ARM.
  - e.g. libVLC (https://github.com/adrg/libvlc-go)
- The mplayer process handling has some bugs, and it's currently easy to get out of sync with the transport controls.
- Deduplicate CORS code into an http.Handler middleware
- inotify listen for changes in the media dir and reindex
- Do cross-compilation for RPi ARM
- Grab album art. Discogs has an API.
- Richer metadata model for music, especially for classical music (movement, piece, opus, recording, performer, conductor, soloist...)
- Websocket communication with what the server is doing.
  - When the song starts, pauses, resumes, stops, or changes
  - When the playback advances by one second

## Getting it running

why would you do that to yourself, but OK

1. [Install Mint](https://mint-lang.com/install).
2. [Install Go](https://go.dev/dl/).

Build the Mint application:
```
$ cd ui
$ mint build
```

Change back to the main directory, then build and start the server, telling it where your music files are:
```
$ make
$ ./stereo --dir="/home/you/Music"
```

You can change the port with `--port` and decide whether or not to reindex on startup with `--reindex`.
