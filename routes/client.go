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
type BodyService struct {
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
			// si en body vienen datos (debe venir el id) se ejecuta la búsqueda por id
			var payload BodyService
			// convierte un tipo bytes (con formato json) a un tipo struct
			err = json.Unmarshal(b, &payload)
			if err != nil {
				// imprime todo el stack de errores
				fmt.Printf("%+v\n", err)
				// Imprime cual fue el primer error
				fmt.Println(errors.Cause(err))
				http.Error(w, err.Error(), 500)
				return
			}
			resp, err := controllers.GetClient(payload.ID)
			if err != nil {
				// imprime todo el stack de errores
				fmt.Printf("%+v\n", err)
				// Imprime cual fue el primer error
				fmt.Println(errors.Cause(err))
				http.Error(w, err.Error(), 400)
				return
			}
			// codifica una estructura a una cadena json y la devuelve como respuesta
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				fmt.Printf("%+v\n", err)
				fmt.Println(errors.Cause(err))
				http.Error(w, err.Error(), 500)
				return
			}
		} else {
			// Si no tiene datos el body, se devuelve un arreglo
			resp, err := controllers.GetClients()
			if err != nil {
				fmt.Printf("%+v\n", err)
				fmt.Println(errors.Cause(err))
				http.Error(w, err.Error(), 400)
				return
			}
			// codifica una estructura a una cadena json y la devuelve como respuesta
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
