// Package main is all that's needed to run my site. It's little more
// than a simple http server that deals with one URL (but it does it well!).
package main

import (
	"github.com/gorilla/handlers"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

// IndexHandler checks if the URL.Path contains a dot (.) as in the case
// of filetypes in the URL.  In this case, it uses a different Mux
// to handle such requests, else it uses the default Mux and executes
// the index template.
func indexHandler(writer http.ResponseWriter, request *http.Request) {

	if strings.Contains(request.URL.Path, ".") {
		myMux.ServeHTTP(writer, request)
	} else {
		err := templates.ExecuteTemplate(writer, "index.html", request)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

var (
	projectRoot string
	templates   *template.Template
	goPath      = os.Getenv("GOPATH")
	myMux       = http.NewServeMux()
)

func main() {
	log.Println("Starting server...")

	if goPath == "" {
		projectRoot = "."
	} else {
		projectRoot = path.Join(goPath, "src", "github.com", "darthlukan", "bct")
	}

	templates = template.Must(template.ParseGlob(projectRoot + "/html/*"))

	myMux.Handle("/", http.FileServer(http.Dir(projectRoot)))

	http.HandleFunc("/", indexHandler)

	log.Println("Server loaded, check localhost:8080")
	log.Println(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
