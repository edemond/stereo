component Item {

  property title : String = ""
  property subtext : String = ""
  property onClick : Function(Html.Event, Promise(Never, Void))

  style item {
    background-color: transparent;
    border: 0;
    border-bottom: 1px solid #{Color:ITEM_BORDER};
    color: white;
    cursor: pointer;
    display: flex;
    font-family: "Inter";
    font-size: #{Font:SIZE_M};
    justify-content: space-between;
    padding: #{Layout:GRID_S} #{Layout:GRID_S};
    text-align: left;
    width: 100%;

    &:hover {
      background-color: #{Color:ITEM_HOVER};
    }

    &:active {
      background-color: #{Color:ACCENT_MEDIUM};
    }
  }

  style title {
    color: #{Color:ITEM_TITLE};
    font-family: "Inter";
    font-size: #{Font:SIZE_M};
    font-weight: #{Font:WEIGHT_NORMAL};
    margin-bottom: 4px;
    -webkit-font-smoothing: antialiased;
  }

  style subtext {
    color: #{Color:ITEM_SUBTEXT};
    font-family: "Inter";
    font-size: #{Font:SIZE_S};
    font-weight: #{Font:WEIGHT_LIGHT};
    -webkit-font-smoothing: antialiased;
  }

  fun render() : Html {
    <button::item onClick={onClick}>
      <div>
        <div::title>
          <{ title }>
        </div>
        <div::subtext>
          <{ subtext }>
        </div>
      </div>
    </button>
  }
}
