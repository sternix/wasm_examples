// +build js,wasm

/*
https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_sort_list_desc
*/
package main

import (
	"github.com/sternix/wasm"
	"sort"
)

var cities = []string{
	"Oslo",
	"Stockholm",
	"Helsinki",
	"Berlin",
	"Rome",
	"Madrid",
	"Zagreb",
	"Amsterdam",
}

func main() {
	doc := wasm.CurrentDocument()

	ulCities := doc.HTMLElementById("ulCities")

	stringSlice := sort.StringSlice(cities)

	updateCities := func() {
		// Clear items
		for ulCities.FirstChild() != nil {
			ulCities.RemoveChild(ulCities.FirstChild())
		}

		for _, city := range stringSlice {
			li := wasm.NewHTMLLIElement()
			li.SetInnerHTML(city)
			ulCities.AppendChild(li)
		}
	}

	var reverse bool
	btnSort := doc.HTMLElementById("btnSort")
	btnSort.OnClick(func(wasm.MouseEvent) {
		if reverse {
			sort.Sort(sort.Reverse(stringSlice))
			reverse = false
		} else {
			sort.Sort(stringSlice)
			reverse = true
		}
		updateCities()
	})

	// populate list
	updateCities()

	wasm.Loop()
}
