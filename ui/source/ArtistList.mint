component ArtistList {
  property artists : Array(Artist)

  fun renderArtist (artist : Artist) : Html {
    <Artist artist={artist} />
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

  fun compareArtists (a : Artist, b : Artist) : Number {
    compareStrings(a.name, b.name)
  }
  
  fun render() : Html {
    <{ Array.map(renderArtist, sortedArtists) }>
  } where {
    sortedArtists = Array.sort(compareArtists, artists)
  }
}
