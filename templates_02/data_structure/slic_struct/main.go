package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("assign.html"))
}
func main() {

	ps1 := Person{
		Name:"ibrahim",
		Age:24,
	}
	ps2 := Person{
		Name:"Hafast",
		Age:15,
	}
	ps3 := Person{
		Name:"Abdulamlik",
		Age:16,
	}
	ps4 := Person{
		Name:"Khadija",
		Age:10,
	}
	ps5 := Person{
		Name:"Zee",
		Age:9,
	}

	people := []Person{
		ps1,
		ps2,
		ps3,
		ps4,
		ps5,
	}
	//about := map[string]string{"name": "Ibrahim", "occupation": "programming", "age": "24"}
	err := tpl.ExecuteTemplate(os.Stdout, "assign.html", people)
	if err != nil {
		log.Fatalln("sorry error occures", err)
	}
}
