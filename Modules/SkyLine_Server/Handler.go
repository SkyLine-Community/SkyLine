package SkyLine_Network

import (
	"fmt"
	"net/http"
)

const (
	ServerStaticFP = "Modules/SkyLine_Server/Static/Home.html"
)

var Paths = map[string]func(writer http.ResponseWriter, req *http.Request){
	"/": func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Modules/SkyLine_Server/Static/Home.html")
	},
	"/Reference/Manuals.html": func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Modules/SkyLine_Server/Static/Reference/Manuals.html")
	},
	"/Reference/REPL.html": func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Modules/SkyLine_Server/Static/Reference/REPL.html")
	},
	"/Reference/CLI.html": func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Modules/SkyLine_Server/Static/Reference/CLI.html")
	},
	"/Reference/Hello.html": func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Modules/SkyLine_Server/Static/Reference/Hello.html")
	},
	"/Reference/logo.ico": func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Modules/SkyLine_Server/Static/logo.ico")
	},
	"/Reference/logo.png": func(writer http.ResponseWriter, req *http.Request) {
		http.ServeFile(writer, req, "Modules/SkyLine_Server/Static/Reference/logo.png")
	},
	"Banner.png": func(writer http.ResponseWriter, req *http.Request) {
		http.ServeFile(writer, req, "Modules/SkyLine_Server/Static/Banner.png")
	},
	"/Reference/About.html": func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Modules/SkyLine_Server/Static/Reference/About.html")
	},
	"/Reference/Variables.html": func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Modules/SkyLine_Server/Static/Reference/Variables.html")
	},
	"/Reference/First.html": func(writer http.ResponseWriter, req *http.Request) {
		http.ServeFile(writer, req, "Modules/SkyLine_Server/Static/Reference/First.html")
	},
	"/Reference/Tutorials.html": func(writer http.ResponseWriter, req *http.Request) {
		http.ServeFile(writer, req, "Modules/SkyLine_Server/Static/Reference/Tutorials.html")
	},
}

func SkyLine_Network_Handler(writer http.ResponseWriter, requeststream *http.Request) {
	requestpath := requeststream.URL.Path
	var newpath string
	switch requeststream.Method {
	case "GET":
		Paths[requestpath](writer, requeststream)
	case "POST":
		if X := requeststream.ParseForm(); X != nil {
			http.ServeFile(writer, requeststream, newpath)
		}
	default:
		fmt.Fprintf(writer, " Unexpected or Unsupported HTTP method for requests")
	}
}
