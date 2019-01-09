// +build js,wasm

/*
https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_image_compare
*/
package main

import (
	"fmt"
	"github.com/sternix/wasm"
)

var (
	win = wasm.CurrentWindow()
	doc = win.Document()
)

func main() {
	for _, img := range doc.HTMLElementsByClassName("img-comp-overlay") {
		CompareImages(img)
	}
	wasm.Loop()
}

func CompareImages(img wasm.HTMLElement) {
	var (
		clicked bool
		w       int
		h       int
		slider  wasm.HTMLDivElement
	)

	w = img.OffsetWidth()
	h = img.OffsetHeight()
	img.Style().SetProperty("width", fmt.Sprintf("%dpx", w/2))

	getCursorPos := func(e wasm.MouseEvent) int {
		rect := img.BoundingClientRect()
		x := e.PageX() - rect.Left()
		return int(x - win.PageXOffset())
	}

	slide := func(pos int) {
		img.Style().SetProperty("width", fmt.Sprintf("%dpx", pos))
		slider.Style().SetProperty("left", fmt.Sprintf("%dpx", img.OffsetWidth()-(slider.OffsetWidth()/2)))
	}

	slideMove := func(e wasm.MouseEvent) {
		if clicked {
			pos := getCursorPos(e)
			if pos < 0 {
				pos = 0
			}
			if pos > w {
				pos = w
			}
			slide(pos)
		}
	}

	slideReady := func(e wasm.MouseEvent) {
		e.PreventDefault()
		clicked = true
		win.OnMouseMove(slideMove)
	}

	slideFinish := func(wasm.MouseEvent) {
		clicked = false
	}

	slider = wasm.NewHTMLDivElement()
	slider.SetClassName("img-comp-slider")
	img.ParentElement().InsertBefore(slider, img)
	slider.Style().SetProperty("top", fmt.Sprintf("%dpx", (h/2)-(slider.OffsetHeight()/2)))
	slider.Style().SetProperty("left", fmt.Sprintf("%dpx", (w/2)-(slider.OffsetWidth()/2)))
	slider.OnMouseDown(slideReady)

	win.OnMouseUp(slideFinish)
}
