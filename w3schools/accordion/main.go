// +build wasm,js

//https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_accordion_symbol
package main

import (
	"fmt"
	"github.com/sternix/wasm"
)

func main() {
	doc := wasm.CurrentDocument()
	acc := doc.ElementsByClassName("accordion")

	for i := 0; i < len(acc); i++ {
		btn := acc[i].(wasm.HTMLButtonElement)
		btn.OnClick(func(wasm.MouseEvent) {
			this := btn
			this.ClassList().Toggle("active")
			panel := this.NextElementSibling().(wasm.HTMLDivElement)
			if panel.Style().MaxHeight() != "" {
				panel.Style().SetMaxHeight("")
			} else {
				panel.Style().SetMaxHeight(fmt.Sprintf("%dpx", panel.ScrollHeight()))
			}
		})
	}

	wasm.Loop()
}
