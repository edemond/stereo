component AlbumList {
  property albums : Array(Album) = []

  fun getSongCount(album : Album) : String {
    if (count == 1) {
      Number.toString(count) + " song"
    } else {
      Number.toString(count) + " songs"
    }
  } where {
    count = Array.size(album.songs) 
  }

  fun onAlbumClick(event : Html.Event) : Promise(Never, Void) {
    Promise.never()
  }

  fun renderAlbum (album : Album) : Html {
    <Item 
      title={album.title}
      subtext={"Album by #{album.artist} · #{album.year} · #{getSongCount(album)}"}
      onClick={onAlbumClick}
    />
  }

  fun render : Html {
    <{ Array.map(renderAlbum, albums) }>
  }
}
