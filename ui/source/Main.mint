routes {
  / {
    Store.setMode(Mode::Songs)
  }

  /songs {
    Store.setMode(Mode::Songs)
  }

  /albums {
    Store.setMode(Mode::Albums)
  }

  /artists {
    Store.setMode(Mode::Artists)
  }

  /artist/:id (id: Number) {
    Store.setMode(Mode::Albums)
  }
}

enum Mode {
  Albums
  Artist(String) // artist ID
  Artists
  Songs
  Error(String)
}

enum Playback {
  Playing(Song)
  Paused(Song)
  Stopped
}

enum Message {
  PlaybackChanged(Playback)
  TimeChanged(Number) // TODO: Maybe make this part of playback state.
  Unknown(String)
}

component Main {
  connect Store exposing { mode, songs, playback, albums, artists, currentTime }

  use Provider.WebSocket {
    onClose = () : Promise(Never, Void) {
      Promise.never()
    },
    onError = () : Promise(Never, Void) {
      // TODO: Should be more tolerant of websocket errors.
      Store.setError("WebSocket error")
    },
    onMessage = (message : String) : Promise(Never, Void) {
      Store.onWebSocketMessage(message)
    },
    onOpen = (socket : WebSocket) : Promise(Never, Void) {
      Promise.never()
    },
    reconnectOnClose = false,
    url = "ws://localhost:8080/api/ws"
  }

  fun componentDidMount : Promise(Never, Void) {
    parallel {
      songs = Server.getSongs()
      artists = Server.getArtists()
      nowPlayingSong = Server.getNowPlayingSong()
    } then {
      Store.setInitialState(songs, artists, nowPlayingSong)
    } catch String => err {
      Store.setError(err)
    }
  }

  style app {
    @font-face {
      font-family: "Inter";
      src: url("/Inter.var.woff2") format("woff2");
    }

    background-color: #{Color:BLACK};
    color: white;
    font-family: "Inter";
    min-height: 100vh;
  }

  fun render : Html {
    <div::app>
      <Nav />
      <{ page }>
      <{ nowPlaying }>
    </div>
  } where {
    page =
      case (mode) {
        Mode::Albums => <AlbumList albums={albums} />
        Mode::Artist(a) => <ArtistView artistID={a} />
        Mode::Artists => <ArtistList artists={artists} />
        Mode::Songs => <SongList songs={songs}/>
        Mode::Error(str) => <Error message={str} />
      }
    nowPlaying =
      case (playback) {
        Playback::Playing(song) => <NowPlaying song={song} time={currentTime} />
        Playback::Paused(song) => <NowPlaying song={song} time={currentTime} />
        Playback::Stopped => <div />
      }
  }
}
