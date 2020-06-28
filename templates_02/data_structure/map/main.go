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
about :=map[string]string{"name":"Ibrahim","occupation":"programming","age":"24"}
	err := tpl.ExecuteTemplate(os.Stdout, "assign.html", about)
	if err != nil {
		log.Fatalln("sorry error occures", err)
	}
}
