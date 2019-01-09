// +build js,wasm

/*
https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_image_zoom
*/
package main

import (
	"fmt"
	"github.com/sternix/wasm"
)

type Pos struct {
	X int
	Y int
}

var (
	doc wasm.Document
	win wasm.Window
)

func main() {
	win = wasm.CurrentWindow()
	doc = win.Document()

	ImageZoom("myimage", "myresult")

	wasm.Loop()
}

func ImageZoom(imgId, resultId string) {
	img := doc.ElementById(imgId).(wasm.HTMLImageElement)
	result := doc.HTMLElementById(resultId)

	lens := wasm.NewHTMLDivElement()
	lens.SetClassName("img-zoom-lens")
	img.ParentElement().InsertBefore(lens, img)

	cx := result.OffsetWidth() / lens.OffsetWidth()
	cy := result.OffsetHeight() / lens.OffsetHeight()

	result.Style().SetBackgroundImage(fmt.Sprintf("url('%s')", img.Src()))
	result.Style().SetBackgroundSize(fmt.Sprintf("%dpx %dpx", img.Width()*cx, img.Height()*cy))

	getCursorPos := func(e wasm.PointerEvent) Pos {
		a := img.BoundingClientRect()
		x := e.PageX() - a.Left()
		y := e.PageY() - a.Top()

		x = x - win.PageXOffset()
		y = y - win.PageYOffset()

		return Pos{
			X: int(x),
			Y: int(y),
		}
	}

	moveLens := func(e wasm.PointerEvent) {
		e.PreventDefault()
		pos := getCursorPos(e)
		x := pos.X - lens.OffsetWidth()/2
		y := pos.Y - lens.OffsetHeight()/2

		if x > (img.Width() - lens.OffsetWidth()) {
			x = img.Width() - lens.OffsetWidth()
		}

		if x < 0 {
			x = 0
		}

		if y > (img.Height() - lens.OffsetHeight()) {
			y = img.Height() - lens.OffsetHeight()
		}

		if y < 0 {
			y = 0
		}

		lens.Style().SetLeft(fmt.Sprintf("%dpx", x))
		lens.Style().SetTop(fmt.Sprintf("%dpx", y))
		result.Style().SetBackgroundPosition(fmt.Sprintf("-%dpx -%dpx", (x*cx), (y*cy)))
	}

	lens.OnPointerMove(func(e wasm.PointerEvent) {
		moveLens(e)
	})

	img.OnPointerMove(func(e wasm.PointerEvent) {
		moveLens(e)
	})
}
