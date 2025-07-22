package routes

import (
	"net/http"
	"net/url"

	"artweb/helpers"
	"artweb/model"
)

// handle the "/" route
func ArtWebHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found ", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodPost {
		input := r.FormValue("input")
		banner := r.FormValue("banner")
		http.Redirect(w, r, "/artstyle?input="+url.QueryEscape(input)+"&banner="+url.QueryEscape(banner), http.StatusSeeOther)
		return
	}

	helpers.ParseAndExecute("templates/index.html", model.Data{}, 200, w, false)
}
