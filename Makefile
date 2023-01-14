UI_SRC_FILES=ui/source/Main.mint
UI_BUILD_FILES=ui/dist/index.js

all: ui/dist/index.js
	go build

ui/dist/index.js: $(UI_SRC_FILES)
	cd ui && mint build

pi:
	GOOS=linux GOARCH=arm go build -o stereo-arm
