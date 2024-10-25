package main

import (
	"fmt"
	"os"
	"github.com/Douirat/ascii-art-web/logic"
	"html/template"
)

var Tmpl *template.Template

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("The number of argumensts is incorrect")
		return
	}
	data, _ := logic.Readfile("../data/thinkertoy.txt")
	if len(data) == 0 {
		fmt.Println("file data is not valid")
		return
	}
	println(len(data))
	result := logic.InputHandler(args[0], data)
	if len(result) == 0 {
		fmt.Println("No returned result")
		return
	}
	fmt.Println(result)
}