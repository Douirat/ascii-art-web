package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"github.com/Douirat/ascii-art-web/routers"
	"github.com/Douirat/ascii-art-web/controllers"
	"github.com/gorilla/mux"
)



func main() {
	router := mux.NewRouter()
	var err error
	controllers.Tmpl, err = template.ParseGlob("../views/*.html")
	if err != nil {
		log.Fatal("Template parsing error: ", err)
		return
	} 
	routers.Routing(router)
	fmt.Println("Server starts at port 9090...")
	err = http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("Error listening to tcp connection", err)
	}
}
