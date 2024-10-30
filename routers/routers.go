package routers

import(
	"github.com/Douirat/ascii-art-web/controllers"
	"github.com/gorilla/mux"
)

func Routing(router *mux.Router) {
	router.HandleFunc("/home", controllers.ProcessRequests).Methods("GET")
	router.HandleFunc("/ascii-art", controllers.RenderResult).Methods("POST")
}