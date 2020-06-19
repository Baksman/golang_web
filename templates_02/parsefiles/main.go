package main

import (
	"log"
	"os"
	"text/template"
)

//Parse the execute
func main() {
	tpl, err := template.ParseFiles("first.gmao")
	if err != nil {
		log.Fatalln("erro parsing first file", err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("couldnot print out file", err)
	}
	tpl, err = tpl.ParseFiles("second.gmao", "third.gmao")
	if err != nil {
		log.Fatalln("Error occureed while", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "second.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("third.gmao", "third.gmao")
	if err != nil {
		log.Fatalln("Error occureed while", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "third.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
