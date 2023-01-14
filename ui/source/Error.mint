component Error {
  property message : String = ""

  style error {
    background-color: pink;
    box-sizing: border-box;
    color: black;
    height: 100vh;
    padding: 40px;
    position: fixed;
    top: 0;
    width: 100vw;
  }

  style speech {
    display: inline-block;
    transform-origin: left center;
    transform: rotateZ(-10deg);
  }

  fun render : Html {
    <div::error>
      <h1>"(\\_l) "<span::speech>"<( bah )"</span></h1>
      <{message}>
    </div>
  }
}
