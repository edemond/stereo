component Menu {

  property isOpen : Bool = false

  style menu {
    background-color: #{Color:BLACK};
    display: none;
    flex-direction: column;
    height: 0;
    opacity: 0;
    transition: height 0ms, opacity 100ms ease-out, visibility 100ms linear, transform 150ms ease-out;
    transform: translateY(10px);
    width: 100%;

    if (isOpen) {
      display: flex;
      height: 100vh;
      opacity: 1;
      transform: translateY(0);
    }
  }

  style link {
    align-items: center;
    border-bottom: 1px solid #{Color:MENU_ITEM_BORDER};
    color: #{Color:NAV_LINKS};
    display: flex;
    flex-direction: row;
    font-size: #{Font:SIZE_L};
    font-weight: #{Font:WEIGHT_MEDIUM};
    margin: 0;
    padding: 20px 25px;
    vertical-align: middle;
    text-decoration: none;
    user-select: none;

    &:hover {
      background-color: #{Color:GRAY_DARK};
    }

    &:active {
      background-color: white;
      color: black;
    }

    & > svg {
      fill: white;
      margin-right: 10px;
      max-width: 40px;
      min-width: 40px;
    }

    &:active > svg {
      fill: black;
    }
  }
  
  fun render() : Html {
    <div::menu>
      <a::link href="/songs">
        <IconSong />
        "Songs"
      </a>

      <a::link href="/artists">
        <IconBand />
        "Artists"
      </a>

      if (Flags:SHOW_ALBUMS) {
        <a::link href="/albums">
          <IconCD />
          "Albums"
        </a>
      }
    </div>
  }
}
