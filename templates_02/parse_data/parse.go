package main

import (
	"log"
	"os"
	"text/template"

)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("parse.html"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "parse.html", 100)
	if err != nil {
		log.Fatalln("sorry error occures", err)
	}
}
