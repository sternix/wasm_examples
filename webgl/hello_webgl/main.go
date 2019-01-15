// +build js,wasm

/*
https://jameshfisher.com/2017/09/27/webgl-hello-world.html
*/
package main

import (
	"fmt"
	"github.com/sternix/wasm"
)

func main() {
	canvasEl := wasm.CurrentDocument().ElementById("canvas").(wasm.HTMLCanvasElement)
	ctx := canvasEl.ContextWebGL()
	if ctx == nil {
		fmt.Println("Your browser does not support WebGL")
		return
	}

	ctx.ClearColor(0, 1, 0, 1)
	ctx.Clear(wasm.COLOR_BUFFER_BIT)
}
