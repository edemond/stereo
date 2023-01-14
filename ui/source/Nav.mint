component Nav {
  connect Store exposing { menuOpen }

  const MENU_ITEMS_HEIGHT = "50px"

  style nav {
    position: fixed;
    top: 0;
    width: 100%;
  }

  style header {
    -webkit-backdrop-filter: blur(5px);
    align-items: center;
    backdrop-filter: blur(5px);
    background-color: #{Color:NAV_BACKGROUND};
    border-top: 1px solid #{Color:NAV_BORDER};
    box-sizing: border-box;
    display: flex;
    flex-direction: row;
    font-family: "Inter";
    justify-content: space-between;
    padding: 10px;
  }
  
  style spacer {
    background: red;
    height: 70px;
  }

  style searchBox {
    background-color: #{Color:SEARCH_BOX_BACKGROUND};
    border: 1px solid #{Color:SEARCH_BOX_BORDER};
    box-sizing: border-box;
    border-radius: 4px;
    color: #{Color:SEARCH_BOX_COLOR};
    font-family: "Inter";
    font-size: #{Font:SIZE_M};
    font-weight: #{Font:WEIGHT_MEDIUM};
    height: #{MENU_ITEMS_HEIGHT};
    margin-right: 10px;
    padding: 5px 15px;
    width: 100%;
  }

  style menuButton {
    background-color: transparent;
    border: none;
    color: #{Color:NAV_MENU_BUTTON};
    cursor: pointer;
    height: #{MENU_ITEMS_HEIGHT};
    
    if (menuOpen) {
      transform: rotate(45deg);
    }
  }

  fun onMenuButtonClick : Promise(Never, Void) {
    Store.setMenuOpen(!menuOpen)
  }

  fun renderButton : Html {
    if (menuOpen) {
      <button::menuButton onClick={onMenuButtonClick}>
        <svg width="40px" height="40px" viewBox="0 0 100 100">
          <rect fill="#{Color:GRAY_MEDIUM}" x="0" y="46" width="100" height="8" />
          <rect fill="#{Color:GRAY_MEDIUM}" x="46" y="0" height="100" width="8" />
        </svg>
      </button>
    } else {
      <button::menuButton onClick={onMenuButtonClick}>
        <svg width="40px" height="40px" viewBox="0 0 100 100">
          <rect fill="#{Color:GRAY_MEDIUM}" x="5" y="15" width="90" height="#{barHeight}" />
          <rect fill="#{Color:GRAY_MEDIUM}" x="5" y="#{50 - (barHeight / 2)}" width="90" height="#{barHeight}" />
          <rect fill="#{Color:GRAY_MEDIUM}" x="5" y="#{85 - barHeight}" width="90" height="#{barHeight}" />
        </svg>
      </button>
    }
  } where {
    barHeight = 7
  }

  fun render : Html {
    <div>
      <div::spacer />
      <div::nav>
        <header::header>
          if (Flags:SHOW_SEARCH_BAR) {
            <input::searchBox placeholder="Search..."></input>
          } else {
            <div />
          }
          <{renderButton()}>
        </header>
        <Menu isOpen={menuOpen} />
      </div>
    </div>

  }
}
