// +build js,wasm

/*
https://www.w3schools.com/howto/tryit.asp?filename=tryhow_html_include_2
*/
package main

import (
	"fmt"
	"github.com/sternix/wasm"
	"io/ioutil"
	"net/http"
	"path"
)

var win = wasm.CurrentWindow()

func main() {
	IncludeHTML()

	wasm.Loop()
}

func IncludeHTML() {
	doc := win.Document()
	loc := path.Dir(win.Document().URL())

	elements := doc.ElementsByTagName("*")
	for _, elm := range elements {
		file := elm.Attribute("w3-include-html")
		// interesting null values returns as "null"
		if file != "" && file != "null" {
			file = fmt.Sprintf("%s/%s", loc, file)
			he := elm.(wasm.HTMLElement)
			resp, err := http.Get(file)
			if err != nil {
				he.SetInnerHTML(err.Error())
				continue
			}

			defer resp.Body.Close()
			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				he.SetInnerHTML(err.Error())
				continue
			}
			he.SetInnerHTML(string(content))
			he.RemoveAttribute("w3-include-html")
		}
	}
}
