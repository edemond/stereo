component Eyebrow {
  property children : Array(Html) = []

  style eyebrow {
    color: #{Color:EYEBROW_COLOR};
    font-family: "Inter";
    font-size: #{Font:SIZE_S};
    font-weight: #{Font:WEIGHT_MEDIUM};
    margin: #{Layout:GRID_S};
    margin-top: #{Layout:GRID_L};
  }

  fun render() : Html {
    <p::eyebrow>
      <{ children }>
    </p>
  }
}
