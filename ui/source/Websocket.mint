module Websocket {

  fun parse(message : String) : Promise(Never, Message) {
    if (Array.size(tokens) <= 0) {
      Promise.resolve(Message::Unknown(message))
    } else {
      parseMessage(message, tokens)
    }
  } where {
    tokens = String.split(" ", message)
  }

  fun parseMessage(message : String, tokens : Array(String)) : Promise(Never, Message) {
    case (token) {
      "play" => parsePlay(message, tokens)
      "pause" => parsePause(message, tokens)
      "stop" => Promise.resolve(Message::PlaybackChanged(Playback::Stopped))
      "time" => parseTime(message, tokens)
      => Promise.resolve(Message::Unknown(message))
    }
  } where {
    token = Maybe.withDefault("", tokens[0])
  }

  // play 1
  fun parsePlay(message : String, tokens : Array(String)) : Promise(Never, Message) {
    sequence {
      songId = Maybe.toResult("No song ID parameter", tokens[1])
      song = Maybe.toResult("Couldn't find song", Store.getSongByID(songId))
      Message::PlaybackChanged(Playback::Playing(song))
    } catch String => err {
      Message::Unknown(message)
    }
  }

  // time 123
  fun parseTime(message : String, tokens : Array(String)) : Promise(Never, Message) {
    sequence {
      token = Maybe.toResult("No timestamp parameter", tokens[1])
      time = Maybe.toResult("Timestamp not a number", Number.fromString(token))
      Message::TimeChanged(time)
    } catch String => err {
      Message::Unknown(message)
    }
  }

  // pause 1
  fun parsePause(message : String, tokens : Array(String)) : Promise(Never, Message) {
    sequence {
      songId = Maybe.toResult("No song ID parameter", tokens[1])
      song = Maybe.toResult("Couldn't find song", Store.getSongByID(songId))
      Message::PlaybackChanged(Playback::Paused(song))
    } catch String => err {
      Message::Unknown(message)
    }
  }
}
