#!/bin/sh

REPOS='
github.com/gorilla/websocket
github.com/go-gl/mathgl/mgl32
github.com/ByteArena/box2d
github.com/lucasb-eyer/go-colorful
'

for repo in $REPOS; do
	go get -u -v $repo
done


WASM_REPOS='
github.com/sternix/wasm
'

export GOARCH=wasm GOOS=js

for repo in $WASM_REPOS; do
	go get -u -v $repo
done
