// +build js,wasm

/*
https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_filter_elements
*/
package main

import (
	"fmt"
	"github.com/sternix/wasm"
	"strings"
)

var (
	doc wasm.Document
)

func main() {
	doc = wasm.CurrentDocument()

	btnAll := doc.ElementById("btnAll").(wasm.HTMLButtonElement)
	btnAll.OnClick(func(wasm.MouseEvent) {
		filterSelection("all")
		highlightActive(btnAll)
	})

	btnCars := doc.ElementById("btnCars").(wasm.HTMLButtonElement)
	btnCars.OnClick(func(wasm.MouseEvent) {
		filterSelection("cars")
		highlightActive(btnCars)
	})

	btnAnimals := doc.ElementById("btnAnimals").(wasm.HTMLButtonElement)
	btnAnimals.OnClick(func(wasm.MouseEvent) {
		filterSelection("animals")
		highlightActive(btnAnimals)
	})

	btnFruits := doc.ElementById("btnFruits").(wasm.HTMLButtonElement)
	btnFruits.OnClick(func(wasm.MouseEvent) {
		filterSelection("fruits")
		highlightActive(btnFruits)
	})

	btnColors := doc.ElementById("btnColors").(wasm.HTMLButtonElement)
	btnColors.OnClick(func(wasm.MouseEvent) {
		filterSelection("colors")
		highlightActive(btnColors)
	})

	// TODO
	// panic: JavaScript error: Argument 1 of EventTarget.dispatchEvent does not implement interface Event.
	// eh.Dispatch()

	filterSelection("all")
	highlightActive(btnAll)

	wasm.Loop()
}

func highlightActive(btn wasm.HTMLButtonElement) {
	current := doc.ElementsByClassName("active")
	if len(current) > 0 {
		current[0].SetClassName(strings.ReplaceAll(current[0].ClassName(), " active", ""))
		btn.SetClassName(btn.ClassName() + " active")
	}
}

func filterSelection(filter string) {
	if filter == "all" {
		filter = ""
	}
	for _, item := range doc.ElementsByClassName("filterDiv") {
		removeClass(item, "show")
		if strings.Index(item.ClassName(), filter) > -1 {
			addClass(item, "show")
		}
	}
}

func isExist(slc []string, item string) bool {
	for _, str := range slc {
		if str == item {
			return true
		}
	}
	return false
}

func addClass(elm wasm.Element, name string) {
	cnames := strings.Split(elm.ClassName(), " ")
	cls := elm.ClassName()

	for _, cname := range strings.Split(name, " ") {
		if !isExist(cnames, cname) {
			cls += fmt.Sprintf(" %s", cname)
		}
	}
	if cls != elm.ClassName() {
		elm.SetClassName(cls)
	}
}

func removeClass(elm wasm.Element, name string) {
	cnames := strings.Split(elm.ClassName(), " ")
	for _, cname := range strings.Split(name, " ") {
		for i := 0; i < len(cnames); i++ {
			if cname == cnames[i] {
				// poor mans's splice
				copy(cnames[i:], cnames[i+1:])
				cnames[len(cnames)-1] = ""
				cnames = cnames[:len(cnames)-1]
			}
		}
	}
	elm.SetClassName(strings.Join(cnames, " "))
}
