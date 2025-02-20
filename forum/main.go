package forum

import "net/http"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler()
	})
	http.ListenAndServe(":8080", nil)
}