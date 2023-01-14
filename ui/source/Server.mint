module Server {

  fun getArtists() : Promise(String, Array(Artist)) {
    sequence {
      response = Http.get("/api/artists")
        |> Http.send()

      body = Json.parse(response.body)
        |> Maybe.toResult("JSON parse error in artists")

      artists = decode body as Array(Artist)
      Promise.resolve(artists)
    } catch Http.ErrorResponse => error {
      Promise.reject("Couldn't fetch artists: #{httpRequestError(error.type)}")
    } catch Object.Error => error {
      Promise.reject("Error decoding artists JSON to Array(String): " + Object.Error.toString(error)) 
    } catch String => error {
      Promise.reject(error)
    }
  }

  fun getSongs() : Promise(String, Array(Song)) {
    sequence {
      response = Http.get("/api/songs")
        |> Http.send()

      body = Json.parse(response.body)
        |> Maybe.toResult("JSON parse error in songs")

      songs = decode body as Array(Song)
      Promise.resolve(songs)
    } catch Http.ErrorResponse => error {
      Promise.reject("Couldn't fetch songs: #{httpRequestError(error.type)}")
    } catch Object.Error => error {
      Promise.reject("Error decoding JSON to Array(Song): " + Object.Error.toString(error)) 
    } catch String => error {
      Promise.reject(error)
    }
  }

  fun getNowPlayingSong() : Promise(String, Maybe(Song)) {
    sequence {
      response = Http.get("/api/nowplaying")
        |> Http.send()

      body = Json.parse(response.body)
        |> Maybe.toResult("JSON parse error in nowplaying song")

      song = decode body as Maybe(Song)
      Promise.resolve(song)
    } catch Http.ErrorResponse => error {
      Promise.reject("Couldn't fetch now-playing song: #{httpRequestError(error.type)}")
    } catch Object.Error => error {
      Promise.reject("Error decoding JSON to Song: " + Object.Error.toString(error)) 
    } catch String => error {
      Promise.reject(error)
    }
  }

  fun playSong(id : String) : Promise(String, Maybe(Song)) {
    sequence {
      response = Http.post("/api/play?id=#{id}")
        |> Http.send()

      body = Json.parse(response.body)
        |> Maybe.toResult("JSON parse error in playSong response")

      song = decode body as Maybe(Song)
      
      Promise.resolve(song)
    } catch Http.ErrorResponse => error {
      Promise.reject(httpRequestError(error.type))
    } catch Object.Error => error {
      Promise.reject("Error decoding JSON to Song: " + Object.Error.toString(error)) 
    } catch String => error {
      Promise.reject("Error playing song: #{error}")
    }
  }

  fun stop() : Promise(String, Void) {
    sequence {
      Http.post("/api/stop")
        |> Http.send()
      Promise.resolve(void)
    } catch Http.ErrorResponse => error {
      Promise.reject(httpRequestError(error.type))
    }
  }

  fun httpRequestError (error : Http.Error) : String {
    case (error) {
      Http.Error::Aborted => "HTTP request aborted"
      Http.Error::BadUrl => "Bad URL"
      Http.Error::NetworkError => "Network error"
      Http.Error::Timeout => "Timeout"
    }
  }
}
