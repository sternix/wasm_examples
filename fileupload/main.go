// +build js,wasm

/*
implementation html5 fileupload example in Go and Wasm

reference example is
https://www.sitepoint.com/html5-javascript-file-upload-progress-bar/
*/

package main

import (
	"bytes"
	"fmt"
	"github.com/sternix/wasm"
	"net/http"
	"strings"
)

var (
	messages wasm.HTMLDivElement
	progress wasm.HTMLDivElement
	doc      wasm.Document
)

func Log(msg string) {
	messages.SetInnerHTML(msg + messages.InnerHTML())
}

func fileDragHover(e wasm.Event) {
	e.StopPropagation()
	e.PreventDefault()

	if elm, ok := e.Target().(wasm.HTMLElement); ok {
		clsName := ""
		if e.Type() == "dragover" {
			clsName = "hover"
		}
		elm.SetClassName(clsName)
	}
}

func fileSelectHandler(e wasm.Event) {
	fileDragHover(e)

	var files []wasm.File
	if de, ok := e.(wasm.DragEvent); ok {
		files = de.DataTransfer().Files()
	} else if ie, ok := e.Target().(wasm.HTMLInputElement); ok {
		files = ie.Files()
	}

	for _, f := range files {
		parseFile(f)
		uploadFile(f)
	}
}

func parseFile(file wasm.File) {
	msg := fmt.Sprintf(`
		<p>File information:
		Name: <strong>%s</strong>
		Type: <strong>%s</strong>
		Size: <strong>%d</strong>
		</p>
	`, file.Name(), file.Type(), file.Size())
	Log(msg)

	if strings.Contains(file.Type(), "image") {
		reader := wasm.NewFileReader()
		reader.OnLoadEnd(func(e wasm.ProgressEvent) {
			msg := fmt.Sprintf(`
				<p>
					<strong>%s</strong><br/>
					<img src="%s"/>
				</p>
			`, file.Name(), string(reader.Result()))
			Log(msg)
		})
		reader.ReadAsDataURL(file)
	} else if strings.Contains(file.Type(), "text") {
		reader := wasm.NewFileReader()
		reader.OnLoadEnd(func(e wasm.ProgressEvent) {
			result := string(reader.Result())
			result = strings.ReplaceAll(result, "<", "&lt;")
			result = strings.ReplaceAll(result, ">", "&gt;")
			msg := fmt.Sprintf(`
				<p>
					<strong>%s</strong>
				</p>
				<pre>%s</pre>
			`, file.Name(), result)
			Log(msg)
		})
		reader.ReadAsText(file)
	}
}

func uploadFile(file wasm.File) {
	reader := wasm.NewFileReader()

	child := wasm.NewHTMLParagraphElement()
	child.AppendChild(doc.CreateTextNode("upload " + file.Name()))
	progress.AppendChild(child)

	reader.OnProgress(func(e wasm.ProgressEvent) {
		pc := 100 - int(e.Loaded()/e.Total()*100)
		child.Style().SetBackgroundPosition(fmt.Sprintf("%d% 0", pc))
	})

	reader.OnLoadEnd(func(e wasm.ProgressEvent) {
		go func() {
			req, err := http.NewRequest("POST", "/upload", bytes.NewBuffer(reader.Result()))
			if err != nil {
				fmt.Println(err)
				child.SetClassName("failure")
				return
			}
			req.Header.Set("FILENAME", file.Name())

			client := &http.Client{}
			resp, err := client.Do(req)
			defer resp.Body.Close()
			if err != nil {
				fmt.Println(err)
				child.SetClassName("failure")
			} else {
				child.SetClassName("success")
			}
		}()
	})
	reader.ReadAsArrayBuffer(file)
}

func main() {
	doc = wasm.CurrentWindow().Document()
	messages = doc.ElementById("messages").(wasm.HTMLDivElement)
	progress = doc.ElementById("progress").(wasm.HTMLDivElement)

	fileSelect := doc.ElementById("fileselect").(wasm.HTMLInputElement)
	fileDrag := doc.ElementById("filedrag").(wasm.HTMLDivElement)

	fileSelect.OnChange(fileSelectHandler)

	fileDrag.On("dragover", fileDragHover)
	fileDrag.On("dragleave", fileDragHover)
	fileDrag.On("drop", fileSelectHandler)

	wasm.Loop()
}
