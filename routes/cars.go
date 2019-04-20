package routes

import (
	"log"
	"net/http"
)

// CarsServeMux is a struct which behaves as a multiplexer and implements ServeHTTP interface
type CarsServeMux struct {
}

// This is the function handler to be overridden
func (c *CarsServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	switch {
	case r.URL.Path == "/cars" && r.Method == "GET":
		log.Println("two")
	case r.URL.Path == "/cars" && r.Method == "POST":
		log.Println("three")
	default:
		http.NotFound(w, r)
	}
	return
}

// CarsRoute function behaves as a constructor for CarsServeMux
func CarsRoute() *CarsServeMux {
	return &CarsServeMux{} // return a ClientServeMux object.
}
