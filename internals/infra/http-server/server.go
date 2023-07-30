package httpserver

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func createHandlers() {

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

		if err != nil {
			panic("Cannot write to httpResponseWriter")
		}
	})
}

func Listen() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	createHandlers()

	s := http.Server{
		Addr: ":" + port,
	}

	fmt.Println("Server is listenig the port " + port)
	log.Fatal(s.ListenAndServe())

}
