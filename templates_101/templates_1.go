package main

import (
	//"fmt"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	//fmt.Println(tpl)
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	str := fmt.Sprint(`
	<!DOCTYPE HTML>
	<html lang = "en"> 
	<head> 
	<meta charset = "utf-8">
	<title> 
	Hello world
	</title>
	<h1>
	My name if  ` + name + `
	</h1>
	</head>
	</html>
	`)
	fmt.Println(str)

	nf, err := os.Create("new.html")
	defer nf.Close()
	if err != nil {
		log.Panic(err)
	}
	io.Copy(nf, strings.NewReader(str))

}