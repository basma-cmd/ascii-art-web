package controllers

import (
	"html/template"
	"net/http"
	"strings"

	"asciiartweb/internal/asciiart"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Output string
		Err    string
	}

	tmpl, err := template.ParseFiles("views/pages/home.html")
	if err != nil {
		http.Error(w, "Error Internal Server", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		input := r.FormValue("text")
		banner := r.FormValue("banner")

		ArtStyle, errAsciiArt := asciiart.AsciiArt(input, banner)
		if errAsciiArt != nil {
			if strings.HasPrefix(errAsciiArt.Error(), "The string contain characters outside the ASCII range...") {
				tmpl.Execute(w, data{Err: "The string contain characters outside the ASCII range..."})
				return
			}
			http.Error(w, "Error Internal Server", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data{Output: ArtStyle})
		return
	}

	tmpl.Execute(w, nil)
}
