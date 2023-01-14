component ArtistView {
  property artistID : String

  style artist {
    color: #{Color:ITEM_TITLE};
    font-family: "Inter";
    font-size: #{Font:SIZE_L};
    font-weight: #{Font:WEIGHT_NORMAL};
    margin: #{Layout:GRID_S};
    -webkit-font-smoothing: antialiased;
  }

  fun renderAlbums(albums : Array(Album)) : Html {
    if (numAlbums <= 0) {
      <></>
    } else {
      <>
        <Eyebrow>
          "ALBUMS (#{numAlbums})"
        </Eyebrow>
        <AlbumList albums={albums} />
      </>
    }
  } where {
    numAlbums = Array.size(albums)
  }

  fun renderSongs(songs : Array(Song)) : Html {
    if (numSongs <= 0) {
      <></>
    } else {
      <>
        <Eyebrow>
          "SONGS (#{numSongs})"
        </Eyebrow>
        <SongList songs={songs} />
      </>
    }
  } where {
    numSongs = Array.size(songs)
  }

  fun renderView(artist : Artist) : Html {
    <div>
      <h1::artist>
        <{artist.name}>
      </h1>
      <{ renderAlbums(albums) }>
      <{ renderSongs(songs) }>
    </div>
  } where {
    albums = Store.getAlbumsByArtistID(artistID)
    songs = Store.getSongsByArtistID(artistID)
  }

  fun render() : Html {
    view
  } where {
    artist = Store.getArtistByID(artistID)
    view = case (Store.getArtistByID(artistID)) {
      Maybe::Just(a) => renderView(a)
      Maybe::Nothing() => <div /> // Should be a loading spinner!
    }
  }
}
