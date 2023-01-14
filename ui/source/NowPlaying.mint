component NowPlaying {

  property song : Song
  property time : Number

  style nowplaying {
    backdrop-filter: blur(5px);
    -webkit-backdrop-filter: blur(5px);
    background-color: #{Color:NOW_PLAYING_BACKGROUND};
    bottom: 0;
    box-shadow: black 0px 0px 30px 30px;
    box-sizing: border-box;
    color: #{Color:NOW_PLAYING_COLOR};
    opacity: 0.98;
    padding: 0 #{Layout:GRID_M};
    position: fixed;
    width: 100%;
  }

  style transport {
    display: flex;
    flex-direction: row;
    justify-content: center;
    margin: 0 auto #{Layout:GRID_M} auto;
    width: 100%;
  }

  style transportButton {
    background: transparent;
    border: none;
    cursor: pointer;
    fill: #{Color:NOW_PLAYING_BUTTON_COLOR};
    margin: 0 #{Layout:GRID_S};
    padding: #{Layout:GRID_S};

    &:hover {
      background: #{Color:NOW_PLAYING_BUTTON_HOVER_BACKGROUND};
      fill: #{Color:NOW_PLAYING_BUTTON_HOVER_COLOR};
    }
  }

  style title {
    font-size: #{Font:SIZE_M};
    font-weight: 600;
    margin-bottom: 0px;
    padding: 0;
    text-align: center;
  }

  style artist {
    color: #{Color:NOW_PLAYING_ARTIST_COLOR};
    font-size: 15px;
    margin: 5px 0;
    text-align: center;
  }

  style progress {
    box-sizing: border-box;
    height: 5px;
    margin: #{Layout:GRID_M} 0;
    width: 100%;

    &[value] {
      -webkit-appearance: none;
      background-color: #{Color:NOW_PLAYING_PROGRESS_BAR_BACKGROUND};
      border: none;
      width: 100%;
    }

    &[value]::-moz-progress-bar {
    }
  }

  style progressArea {
    display: flex;
    flex-direction: row;
    margin: 0 auto;
    max-width: 500px;
  }

  style timestamp {
    color: #{Color:GRAY_LIGHT};
    font-size: 12px;
    margin: 15px 10px;
    vertical-align: middle;
  }

  style left {
    margin-left: 0;
  }
  style right {
    margin-right: 0;
  }

  fun render() : Html {
    <div::nowplaying>
      <div>
        <p::title>
          <{song.title}>
        </p>
        <p::artist>
          <{song.artist}>
        </p>
      </div>
      if (Flags:SHOW_PROGRESS_BAR) {
        <div::progressArea>
          <p::timestamp::left>
            "0:00"
          </p>
          <progress::progress value="70" max="100" />
          <p::timestamp::right>
            "4:36"
          </p>
        </div>
      }
      <div::transport>
        if (Flags:SHOW_PAUSE_BUTTON) {
          <{renderPlayPauseButton()}>
        }
        <{renderStopButton()}>
      </div>
    </div>
  }

  fun onStopButtonClick() : Promise(Never, Void) {
    sequence {
      success = Server.stop()
      Store.setNowPlayingSong(Maybe.nothing())
    } catch String => err {
      Store.setError(err)
    }
  }

  fun renderStopButton() : Html {
    <button::transportButton onClick={onStopButtonClick}>
      <svg width="#{Layout:GRID_M}" height="#{Layout:GRID_M}" viewBox="0 0 100 100">
        <rect x="0" y="0" width="100" height="100" />
      </svg>
    </button>
  }

  fun renderPlayPauseButton() : Html {
    <button::transportButton>
      <svg width="#{Layout:GRID_M}" height="#{Layout:GRID_M}" viewBox="0 0 100 100">
        <rect x="5" y="0" width="35" height="100" />
        <rect x="60" y="0" width="35" height="100" />
      </svg>
    </button>
  }
}
