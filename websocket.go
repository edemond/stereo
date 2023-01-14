package main

import(
  "fmt"
  "github.com/edemond/stereo/log"
  "github.com/edemond/stereo/media"
  "github.com/gorilla/websocket"
  "net/http"
  "sync"
)

var _connections sync.Map

var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
  CheckOrigin: checkOrigin,
}

func checkOrigin(r *http.Request) bool {
  // WebSocket connections are allowed to any host. It's
  // up to the server to verify the origin.
  // TODO: check the origin and return false if it's not us
  return true
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Errorf("Couldn't upgrade Websocket request: %v", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  _connections.Store(conn, true)
  go readLoop(conn)

  // TODO: when an action is taken that affects all clients,
  // like changing song playback, blast it out on all conns

  // TODO: gotta write a little protocol for doing this
}

func broadcast(msg []byte) {
  _connections.Range(func(k, v interface{}) bool {
    conn, ok := k.(*websocket.Conn)
    if !ok {
      return true
    }
    conn.WriteMessage(websocket.TextMessage, msg)
    return true
  })
}

func BroadcastPlay(song *media.SongFile) {
  broadcast([]byte(fmt.Sprintf("play %v", song.Id)))
}

func BroadcastStop() {
  broadcast([]byte("stop"))
}

func readLoop(conn *websocket.Conn) {
  msgType, _, err := conn.ReadMessage()
  log.Infof("websocket message received: %v", msgType)

  if err != nil {
    if websocket.IsCloseError(err, websocket.CloseGoingAway) {
      log.Infof("websocket close received: %v", err)
      _connections.Delete(conn)
    } else if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
      log.Infof("unexpected websocket close received: %v", err)
      _connections.Delete(conn)
    } else {
      log.Errorf("unexpected websocket error received: %v", err)
    }
  } else if msgType == websocket.CloseMessage {
    log.Infof("Client closed websocket")
    _connections.Delete(conn)
  }
}
