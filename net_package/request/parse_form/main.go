package main

import(
	"html/template"
	"log"
	"net/http"
)


var tpl *template.Template
type hotDog int 

func (m hotDog) ServeHTTP(w http.ResponseWriter ,req *http.Request){
		err := req.ParseForm()
		err != nil{
			log.Fatalln(err)
		}

}