package helpers

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"artweb/model"
)

func ParseAndExecute(temp string, data model.Data, status int, w http.ResponseWriter, errNotInASCII bool) {
	if !strings.HasPrefix(data.Error, "The string includes characters outside the ASCII range...") {
		switch status {
		case 400:
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		case 404:
			http.Error(w, "404 Page Not Found", http.StatusNotFound)
			return
		case 500:
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	tmp, err := template.ParseFiles(temp)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	if errNotInASCII {
		data := model.Data{Error: "The string includes characters outside the ASCII range..."}
		w.WriteHeader(status)
		errExec := tmp.Execute(w, data)
		if errExec != nil {
			log.Println(errExec)
		}
		return
	}

	errExec := tmp.Execute(w, data)
	if errExec != nil {
		log.Println(errExec)
	}
}
