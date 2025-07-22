package routes

import (
	"net/http"
	"strings"

	"artweb/app"
	"artweb/helpers"
	"artweb/model"
)

func ArtStyleHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	banner := r.URL.Query().Get("banner")

	ArtStyle, err := app.AsciiArt(input, banner)
	if err != nil {
		if strings.HasPrefix(err.Error(), "The string includes characters outside the ASCII range...") {
			helpers.ParseAndExecute("templates/index.html", model.Data{Error: err.Error()}, http.StatusBadRequest, w, true)
		} else if err.Error() == "input error" {
			helpers.ParseAndExecute("", model.Data{Error: err.Error()}, http.StatusBadRequest, w, false)
		} else {
			helpers.ParseAndExecute("", model.Data{Error: err.Error()}, http.StatusInternalServerError, w, false)
		}
		return
	}
	helpers.ParseAndExecute("templates/artstyle.html", model.Data{Output: ArtStyle}, 200, w, false)
}
