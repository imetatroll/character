package main

import (
	"github.com/gopherjs/gopherjs/js"

	"imetatroll.com/character.git/importer/dndbeyond"
)

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
				obj := js.Global.Get("JSON").Call("parse", fi.Get("target").Get("result"))
				char := &beyond.Character{}
				char.Object = obj

				// println values here when manually testing

				println(char.Data.Name)
				for i, slot := range char.Data.Classes[0].Definition.SpellRules.LevelSpellSlots[1] {
					println(i, slot)
				}
				println(char.Data.Classes[0].Definition.SpellRules.IsRitualSpellCaster)
			})
			fileReader.Call("readAsText", file)
		})
	})
}
