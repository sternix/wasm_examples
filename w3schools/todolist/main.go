// +build js,wasm

/*
https://www.w3schools.com/howto/tryit.asp?filename=tryhow_js_todo
*/
package main

import (
	"github.com/sternix/wasm"
)

type Item struct {
	Subject   string
	Completed bool
}

var (
	input wasm.HTMLInputElement
	myUL  wasm.HTMLElement
	doc   wasm.Document
	items = []Item{
		Item{"Hit the gym", false},
		Item{"Pay bills", true},
		Item{"Meet George", false},
		Item{"Buy eggs", false},
		Item{"Read a book", false},
		Item{"Organize office", false},
	}
)

func main() {
	win := wasm.CurrentWindow()
	doc = win.Document()

	myUL = doc.HTMLElementById("myUL")
	myUL.OnClick(func(e wasm.MouseEvent) {
		elm := e.Target().(wasm.Element)
		if elm.TagName() == "LI" {
			elm.ClassList().Toggle("checked")
		}
	})

	addItem := func() {
		inputValue := input.Value()
		if inputValue == "" {
			win.Alert("You must write something")
		} else {
			addNewItem(Item{inputValue, false})
			input.SetValue("")
		}
	}

	input = doc.ElementById("myInput").(wasm.HTMLInputElement)
	input.OnKeyDown(func(e wasm.KeyboardEvent) {
		if e.Key() == wasm.KeyEnter {
			addItem()
		}
	})

	spanAdd := doc.HTMLElementById("spanAdd")
	spanAdd.OnClick(func(wasm.MouseEvent) {
		addItem()
	})

	for _, item := range items {
		addNewItem(item)
	}

	wasm.Loop()
}

func addNewItem(item Item) {
	li := wasm.NewHTMLLIElement()
	if item.Completed {
		li.SetClassName("checked")
	}
	t := doc.CreateTextNode(item.Subject)
	li.AppendChild(t)
	myUL.AppendChild(li)
	span := wasm.NewHTMLSpanElement()
	txt := doc.CreateTextNode("\u00D7")
	span.SetClassName("close")
	span.AppendChild(txt)
	li.AppendChild(span)

	span.OnClick(func(wasm.MouseEvent) {
		li.Style().SetProperty("display", "none")
	})
}
