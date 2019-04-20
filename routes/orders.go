package routes

import (
	"log"
	"net/http"
)

// OrderServeMux is a struct which behaves as a multiplexer
type OrderServeMux struct {
}

// This is the function handler to be overridden
func (c *OrderServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	switch {
	case r.URL.Path == "/orders" && r.Method == "GET":
		log.Println("GET")
	case r.URL.Path == "/orders" && r.Method == "POST":
		log.Println("POST")
	case r.URL.Path == "/orders" && r.Method == "DELETE":
		log.Println("DELETE")
	case r.URL.Path == "/orders" && r.Method == "PUT":
		log.Println("PUT")
	default:
		http.NotFound(w, r)
	}
	return
}

// OrderRoute function behaves as a constructor for OrderServeMux
func OrderRoute() *OrderServeMux {
	return &OrderServeMux{} // return a OrderServeMux object.
}
