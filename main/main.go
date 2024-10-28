package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Douirat/ascii-art-web/routers"
	"github.com/gorilla/mux"
)



func main() {
	router := mux.NewRouter()
	routers.Routing(router)
	fmt.Println("Server starts at port 9090...")
	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("Error listening to tcp connection", err)
	}
}
