store Store {
  state albums : Array(Album) = Fake.albums()
  state artists : Array(Artist) = []
  state currentTime : Number = 0 // TODO: Consider making this part of Playback.
  state menuOpen : Bool = false
  state mode : Mode = Mode::Songs
  state playback : Playback = Playback::Stopped
  state songs : Array(Song) = []

  fun getAlbumsByArtistID(artistId : String) : Array(Album) {
    Array.select((album : Album) : Bool { album.artistId == artistId }, albums)
  }

  fun getArtistByID(id : String) : Maybe(Artist) {
    Array.find((artist : Artist) : Bool { artist.id == id }, artists)
  }

  fun getPlaybackFromSong(song : Maybe(Song)) : Playback {
    case(song) {
      Maybe::Just(s) => Playback::Playing(s)
      Maybe::Nothing => Playback::Stopped
    }
  }

  fun getSongByID(id : String) : Maybe(Song) {
    Array.find((song : Song) : Bool { song.id == id }, songs)
  }

  fun getSongsByArtistID(artistId : String) : Array(Song) {
    Array.select((song : Song) : Bool { song.artistId == artistId }, songs)
  }

  fun onWebSocketMessage(message : String) : Promise(Never, Void) {
    sequence {
      parsed = Websocket.parse(message)
      case (parsed) {
        Message::PlaybackChanged(p) => next { playback = p }
        Message::TimeChanged(t) => next { currentTime = t }
        Message::Unknown(msg) => setError("Unknown WebSocket message: " + msg)
      }
    }
  }

  fun setError(err : String) : Promise(Never, Void) {
    next {
      mode = Mode::Error(err),
    }
  }

  fun setMenuOpen(isOpen : Bool) : Promise(Never, Void) {
    next { 
      menuOpen = isOpen 
    }
  }

  fun setMode (mode : Mode) : Promise(Never, Void) {
    next { 
      mode = mode,
      menuOpen = false,
    }
  }

  fun setNowPlayingSong(song : Maybe(Song)) : Promise(Never, Void) {
    next {
      playback = getPlaybackFromSong(song),
    }
  }

  // TODO: Kinda don't like this. Isn't there a way to call two setters
  // inside a parallel "then" block without using "next" directly?
  fun setInitialState(
    songs : Array(Song), 
    artists : Array(Artist), 
    nowPlayingSong : Maybe(Song),
  ) : Promise(Never, Void) {
    next {
      songs = songs,
      artists = artists,
      playback = getPlaybackFromSong(nowPlayingSong),
    }
  }

}
