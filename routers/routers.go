package routers

import(
	"github.com/Douirat/ascii-art-web/controllers"
	"github.com/gorilla/mux"
)

func Routing(router *mux.Router) {
	router.HandleFunc("/ascii/", controlers.RenderResult).Methods("GET")
	router.HandleFunc("/", controlers.ProcessRequests).Methods("GET")
}