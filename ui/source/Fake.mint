module Fake {

  fun album (artist : String, title : String) : Album {
    {
      artist = artist,
      artistId = "1",
      title = title,
      songs = [],
      year = 1994
    }
  }

  fun artist (name : String) : Artist {
    {
      id = "1",
      name = name,
    }
  }

  fun albums : Array(Album) {
    [
      Fake.album("Hatfield and the North", "The Rotter's Club"),
      Fake.album("Kali Malone", "The Sacrificial Code"),
      Fake.album("Lupe Fiasco", "Tetsuo and Youth"),
      Fake.album("Smashing Pumpkins", "Siamese Dream"),
      Fake.album("Smashing Pumpkins", "Mellon Collie and the Infinite Sadness"),
      Fake.album("Squarepusher", "Ultravisitor")
    ]
  }

  fun artists : Array(Artist) {
    [
      Fake.artist("Aesop Rock"),
      Fake.artist("Aphex Twin"),
      Fake.artist("Art Tatum"),
      Fake.artist("Meshuggah"),
      Fake.artist("Melvins"),
      Fake.artist("J.S. Bach"),
      Fake.artist("Lupe Fiasco"),
      Fake.artist("Björk"),
      Fake.artist("Squarepusher"),
      Fake.artist("Kali Malone"),
      Fake.artist("SOPHIE"),
      Fake.artist("GFOTY"),
      Fake.artist("John Coltrane"),
      Fake.artist("Kero Kero Bonito"),
      Fake.artist("Kendrick Lamar"),
      Fake.artist("Terror Jr."),
      Fake.artist("Maurizio"),
      Fake.artist("Youth Code"),
      Fake.artist("Emptyset"),
      Fake.artist("Rayner Brown"),
      Fake.artist("Deepchord"),
      Fake.artist("Basic Channel"),
      Fake.artist("Hatfield and the North"),
      Fake.artist("Rhythm & Sound"),
      Fake.artist("Slim Gaillard"),
      Fake.artist("Thou"),
    ]
  }
}
