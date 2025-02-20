package forum

import (
	"net/http"

	"forum/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r)
	})
	http.ListenAndServe(":8080", nil)
}
