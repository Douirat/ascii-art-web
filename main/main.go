package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Douirat/ascii-art-web/logic"
	// "html/template"
)

func RenderResult(wr http.ResponseWriter, rq *http.Request) {
	rq.ParseForm()
	data, _ := logic.Readfile("../data/standard.txt")
	if len(data) == 0 {
		fmt.Println("file data is not valid")
		return
	}
	result := logic.InputHandler("Hello World", data)
	fmt.Fprintf(wr, result)
}

func main() {
	http.HandleFunc("/ascii", RenderResult)
	fmt.Println("Server starts at port 9090...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
