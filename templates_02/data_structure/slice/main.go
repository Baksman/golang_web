package main

import (
	"log"
	"os"
	"text/template"

)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("assign.html"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "assign.html", 100)
	if err != nil {
		log.Fatalln("sorry error occures", err)
	}
}
