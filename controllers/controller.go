package controlers

import (
	"html/template"
	"net/http"

	"github.com/Douirat/ascii-art-web/models"
)

// Declare an object to hold the
type UserInput struct {
	Input, File string
}

var (
	userInput = &UserInput{}
	Tmpl      *template.Template
)

// Process the incomming requests:
func ProcessRequests(wr http.ResponseWriter, rq *http.Request) {
	err := Tmpl.ExecuteTemplate(wr, "../views/index.html", nil)
	if err != nil {
		wr.WriteHeader(http.StatusBadRequest)
	}
}

// Render result to an other web page:
func RenderResult(wr http.ResponseWriter, rq *http.Request) {
	if rq.Method == "POST" {
		userInput.File = rq.FormValue("selectType")
		userInput.Input = rq.FormValue("userInput")
	} else {
		wr.WriteHeader(http.StatusBadRequest)
	}
	if userInput.File != "" && userInput.Input != "" {
		data, _ := models.Readfile(userInput.File)
		if len(data) == 0 {
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
		result := models.InputHandler(userInput.Input, data)
		Tmpl.ExecuteTemplate(wr, "../views.html", result)
	} else {
		wr.WriteHeader(http.StatusBadRequest)
	}
}
