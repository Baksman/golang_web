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
	name := []string{"Baksman", "Ibrahim", "Shehu"}
	err := tpl.ExecuteTemplate(os.Stdout, "assign.html", name)
	if err != nil {
		log.Fatalln("sorry error occures", err)
	}
}
