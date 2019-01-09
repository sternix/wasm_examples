// +build js,wasm

/*
https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_animate_3
*/
package main

import (
	"fmt"
	"github.com/sternix/wasm"
)

func main() {
	var (
		id  int
		pos int
	)

	win := wasm.CurrentWindow()
	doc := win.Document()
	btnMove := doc.HTMLElementById("btnMove")
	elem := doc.HTMLElementById("myAnimation")

	timerCallback := wasm.NewTimerCallback(func(...interface{}) {
		if pos == 350 {
			win.ClearInterval(id)
		} else {
			pos++
			elem.Style().SetTop(fmt.Sprintf("%dpx", pos))
			elem.Style().SetLeft(fmt.Sprintf("%dpx", pos))
		}
	})

	btnMove.OnClick(func(wasm.MouseEvent) {
		pos = 0
		id = win.SetInterval(timerCallback, 10)
	})

	wasm.Loop()
}
