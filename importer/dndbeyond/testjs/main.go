package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	doc := js.Global.Get("document")
	js.Global.Get("window").Set("onload", func() {
		o := doc.Call("getElementById", "CharacterSheet")
		o.Set("onchange", func(evt *js.Object) {
			file := o.Get("files").Index(0)
			if file == nil {
				println("Import file is nil")
				return
			}
			fileReader := js.Global.Get("FileReader").New()
			fileReader.Set("onload", func(fi *js.Object) {
				values := js.Global.Get("JSON").Call("parse", fi.Get("target").Get("result"))
				println("loaded values", values)
			})
			fileReader.Call("readAsText", file)
		})
	})
}
