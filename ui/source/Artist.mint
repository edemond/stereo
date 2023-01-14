record Artist {
  id : String,
  name : String
}

component Artist {
  property artist : Artist

  style artist {
    background-color: transparent;
    color: #{Color:ITEM_TITLE};
    display: block;
    font-family: "Inter";
    font-size: #{Font:SIZE_L};
    font-weight: #{Font:WEIGHT_NORMAL};
    margin: 0 15px;
    padding: 15px 0;
    text-align: left;
    text-decoration: none;
    -webkit-font-smoothing: antialiased;

    &:hover {
      background-color: #{Color:ITEM_HOVER};
    }

    &:active {
      background-color: #{Color:ACCENT_MEDIUM};
    }
  }

  fun render : Html {
    <a::artist href={href}>
      <{artist.name}>
    </a>
  } where {
    href = "/artist/" + artist.id
  }
}
