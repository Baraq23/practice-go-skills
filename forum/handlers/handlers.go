package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

var Tmpl *template.Template


//templates are served correctly everytime
func init() {
	_,filename,_,_ := runtime.Caller(0)
	fmt.Println(filename)
	currentDir := filepath.Dir(filename)
	templatesDir := filepath.Join(currentDir, "../templates")
	Tmpl = template.Must(template.ParseGlob(filepath.Join(templatesDir, "*.html")))
}

//home page handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		ErrorHandler(w, int(http.StatusNotFound), "Sorry we could not find the page you were looking.")
	}
	
}

//handling errors
func ErrorHandler(w http.ResponseWriter, statusCode int, MessageStr string) {
	w.WriteHeader(statusCode)
	data := struct {
		Message string
	}{
		Message: MessageStr,
	}
	err := Tmpl.ExecuteTemplate(w, "error-page.html", data)
	if err != nil {
		log.Println("Error Parsing Error-page.html")
		return
	}
}
