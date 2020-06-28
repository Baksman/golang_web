package main

import (
	"fmt"
	"net/http"
)

type body struct {
}

func (dog body) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Baksman is abig boy yeah becoming cool")
}


func main() {
	var d body 
	http.ListenAndServe(":8080/", d)
}


