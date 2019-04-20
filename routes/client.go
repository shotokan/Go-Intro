package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"raw-service/controllers"

	"github.com/pkg/errors"
)

// Body contains the id
type Body struct {
	ID string `json:"id"`
}

// ClientServeMux is a struct which behaves as a multiplexer
type ClientServeMux struct {
}

// This is the function handler to be overridden
func (c *ClientServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	defer r.Body.Close()
	switch {
	case r.URL.Path == "/clients" && r.Method == "GET":
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		w.Header().Set("content-type", "application/json")
		if len(b) > 0 {
			var payload Body
			err = json.Unmarshal(b, &payload)
			if err != nil {
				fmt.Println(errors.Cause(err))
				http.Error(w, err.Error(), 500)
				return
			}
			resp, err := controllers.GetClient(payload.ID)
			if err != nil {
				fmt.Printf("%+v\n", err)
				fmt.Println(errors.Cause(err))
				http.Error(w, err.Error(), 400)
				return
			}
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				fmt.Printf("%+v\n", err)
				fmt.Println(errors.Cause(err))
				http.Error(w, err.Error(), 500)
				return
			}
		} else {
			resp, err := controllers.GetClients()
			if err != nil {
				fmt.Printf("%+v\n", err)
				fmt.Println(errors.Cause(err))
				http.Error(w, err.Error(), 400)
				return
			}
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				fmt.Printf("%+v\n", err)
				fmt.Println(errors.Cause(err))
				http.Error(w, err.Error(), 500)
				return
			}
		}

	case r.URL.Path == "/clients" && r.Method == "POST":
		log.Println("three")
	default:
		http.NotFound(w, r)
	}
	return
}

// ClientRoute function behaves as a constructor for ClientServeMux
func ClientRoute() *ClientServeMux {
	return &ClientServeMux{} // return a ClientServeMux object.
}
