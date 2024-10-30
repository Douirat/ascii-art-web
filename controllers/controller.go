package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

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
	err := Tmpl.ExecuteTemplate(wr, "index.html", nil)
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
		fmt.Println([]rune(userInput.Input))
		if  IsValidInput(userInput.Input) && IsValidFile(userInput.File) {
			fmt.Printf("The patter is: %s and the input is: %s\n", userInput.File, userInput.Input)
			data, _ := models.Readfile(userInput.File)

			if len(data) == 0 {
				wr.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(wr, "Error: Empty file or banner not found: 500!!!")
				return
			}
			result := models.InputHandler(userInput.Input, data)
			Tmpl.ExecuteTemplate(wr, "index.html", result)
		} else {
			wr.WriteHeader(http.StatusNotFound);
			fmt.Fprintf(wr, "Error: Response not found 404!!!")
		}

	} else {
		wr.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(wr, "Error: bad request 400!!!")
	}
}

// Is user input valid:
func IsValidInput(input string) bool {
	for _, char := range input {
		if char < 32 && char > 127 {
			return false
		}
	}
	fmt.Println("Touch a nerve here!!!")
	return true
}

// Is valid file name:
func IsValidFile(input string) bool {
if strings.EqualFold(input, "../data/shadow.txt") || strings.EqualFold(input, "../data/standard.txt") || strings.EqualFold(input, "../data/thinkertoy.txt") {
	return true
} 
fmt.Println("Touched a nerve")
return false
}
