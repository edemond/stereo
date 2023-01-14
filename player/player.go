package player

import(
)

type Playable interface {
  GetName() string
  GetPath() string
}

type Player interface {
  Play(p Playable)
  Stop()
  OnSongEnd(f func())
}

// TODO: Synchronize access to this.
// Fortunately, it barely matters for now since the max number of concurrent
// users of this server would be the number of people in the house.
var _player Player

func GetPlayer() Player {
  if _player == nil {
    _player = &MPlayer{}
  }
  return _player
}
