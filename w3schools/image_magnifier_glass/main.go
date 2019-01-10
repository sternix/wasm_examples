// +build js,wasm

/*
https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_image_magnifier_glass
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
	win = wasm.CurrentWindow()
	doc = win.Document()
)

func main() {
	Magnify("myimage", 3)
	wasm.Loop()
}

func Magnify(id string, zoom int) {
	img := doc.ElementById(id).(wasm.HTMLImageElement)
	glass := wasm.NewHTMLDivElement()
	glass.SetClassName("img-magnifier-glass")
	img.ParentElement().InsertBefore(glass, img)

	glsStyle := glass.Style()
	glsStyle.SetBackgroundImage(fmt.Sprintf("url('%s')", img.Src()))
	glsStyle.SetBackgroundRepeat("no-repeat")
	glsStyle.SetBackgroundSize(fmt.Sprintf("%dpx %dpx", (img.Width() * zoom), (img.Height() * zoom)))

	w := glass.OffsetWidth() / 2
	h := glass.OffsetHeight() / 2

	bw := 3

	getCursorPos := func(e wasm.MouseEvent) Pos {
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

	moveMagnifier := func(e wasm.MouseEvent) {
		e.PreventDefault()
		pos := getCursorPos(e)
		x := pos.X
		y := pos.Y

		if x > (img.Width() - (w / zoom)) {
			x = img.Width() - (w / zoom)
		}

		if x < (w / zoom) {
			x = w / zoom
		}

		if y > (img.Height() - (h / zoom)) {
			y = img.Height() - (h / zoom)
		}

		if y < (h / zoom) {
			y = h / zoom
		}

		glsStyle.SetLeft(fmt.Sprintf("%dpx", x-w))
		glsStyle.SetTop(fmt.Sprintf("%dpx", y-h))
		glsStyle.SetBackgroundPosition(fmt.Sprintf("-%dpx -%dpx", ((x * zoom) - w + bw), ((y * zoom) - h + bw)))
	}

	glass.OnMouseMove(moveMagnifier)
	img.OnMouseMove(moveMagnifier)
}
