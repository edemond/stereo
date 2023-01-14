package player

import(
  "github.com/edemond/stereo/log"
  "io"
	"os/exec"
  "strings"
  "sync"
)

var _n = 0

// MPlayer implements Player by shelling out to an mplayer installation.
type MPlayer struct {
	cmd *exec.Cmd
  mutex sync.Mutex
  onSongEnd func()
	stdin io.WriteCloser // for sending commands to MPlayer
}

func (m *MPlayer) OnSongEnd(f func()) {
  m.onSongEnd = f
}

// Play plays the given Playable, stopping the current Playable, if any.
func (m *MPlayer) Play(p Playable) {
  _n += 1
  n := _n
	log.Infof("Play: '%v' (%v)\n", p.GetPath(), n)

  m.mutex.Lock()
  defer m.mutex.Unlock()

  if m.cmd != nil {
    m.stopImpl(n)
  }

	var cmd *exec.Cmd
	if isPlaylistType(p.GetPath()) {
		cmd = exec.Command("mplayer", "-playlist", p.GetPath())
	} else {
		cmd = exec.Command("mplayer", p.GetPath())
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Errorf("Error getting stdin of mplayer process: %v\n", err)
		return
	}

	err = cmd.Start()
	if err != nil {
		log.Errorf("Couldn't play '%v': %v\n", p.GetName(), err)
    return
	}

	m.cmd = cmd
  log.Infof("cmd assigned")
	m.stdin = stdin

  log.Infof("Play(%v) waiting for mplayer to terminate...", n)

  // Wait for the song to stop playing.
  go func() {
    m.cmd.Wait()
    log.Infof("cmd.Wait() over, mplayer terminated (%v).", n)

    m.cmd = nil
    log.Infof("cmd nulled out (%v)", n)
    m.onSongEnd()

    log.Infof("Play(%v) done waiting.", n)
  }()
}

func (m *MPlayer) Stop() {
  _n += 1
  n := _n
	log.Infof("Stop called (%v).", n)

  m.mutex.Lock()
  defer m.mutex.Unlock()

  m.stopImpl(n)
  log.Infof("Stop(%v) done.", n)
}

// Do not call without having locked the mutex.
func (m *MPlayer) stopImpl(n int) {
	if m.cmd == nil {
		log.Infof("Can't stop; no song currently playing.")
		return
	}

  // Send mplayer the "q" command to quit.
	written, err := m.stdin.Write([]byte("q"))
	if err != nil {
		log.Infof("Error sending 'q' to mplayer process: %v\n", err)
		return
	}
	if written != 1 {
		log.Infof("Didn't end up writing a 'q' to mplayer stdin: %v\n", err)
		return
	}
  log.Infof("stopImpl(%v) waiting for mplayer to terminate...", n)

  m.cmd.Wait()
  log.Infof("cmd.Wait() over, mplayer terminated (%v).", n)

  m.cmd = nil
  log.Infof("cmd nulled out (%v)", n)
  m.onSongEnd()
}

func isPlaylistType(path string) bool {
	url := strings.ToLower(path)
	return strings.HasSuffix(url, ".pls") || strings.HasSuffix(url, ".m3u");
}
