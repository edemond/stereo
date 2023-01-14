component SongList {
  property songs : Array(Song) = []

  fun onSongClick(song : Song, event : Html.Event) : Promise(a,b) {
    sequence {
      nowPlayingSong = Server.playSong(song.id)
      Store.setNowPlayingSong(nowPlayingSong)
    } catch String => error {
      // TODO: handle error, somehow
      Promise.never()
    }
  }

  fun renderSong (song : Song) : Html {
    <Item 
      title={song.title} 
      subtext={song.artist} 
      onClick={onSongClick(song)}
    />
  }

  fun compareStrings (a : String, b : String) : Number {
    if (a < b) {
      -1
    } else if (a > b) {
      1
    } else {
      0
    }
  }

  fun compareSongs (a : Song, b : Song) : Number {
    if (artistCmp == 0) {
      titleCmp
    } else {
      artistCmp
    }
  } where {
    artistCmp =
      compareStrings(a.artist, b.artist)

    titleCmp =
      compareStrings(a.title, b.title)
  }

  fun render : Html {
    <{ Array.map(renderSong, sortedSongs) }>
  } where {
    sortedSongs =
      Array.sort(compareSongs, songs)
  }
}
