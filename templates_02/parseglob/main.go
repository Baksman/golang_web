package main

import (
	"log"
	"os"
	"text/template"
)

//Parse the execute
func main() {
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln("error parsing first file", err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("couldnot print out file", err)
	}
	tpl, err = tpl.ParseFiles("templates/second.gmao", "third.gmao")
	if err != nil {
		log.Fatalln("Error occureed while", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "templates/second.gmao", nil)
	if err != nil {
		log.Fatalln("Error occureed while", err)
	}

	tpl, err = tpl.ParseFiles("templates/third.gmao", "third.gmao")
	if err != nil {
		log.Fatalln("Error occureed while", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "t   emplates/third.gmao", nil)
	if err != nil {
		log.Fatalln("Error occureed while", err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Error occureed while", err)
	}

}
