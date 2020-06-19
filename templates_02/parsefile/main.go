package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Panic(err)
	}
	nf, err := os.Create("index.hmtl")
	if err != nil {
		log.Panic("eerror creating file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Panic(err)
	}
}
