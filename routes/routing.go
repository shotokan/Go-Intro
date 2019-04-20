package routes

import "net/http"

// LoadRoutes loads all the routes
func LoadRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/clients", ClientRoute())
	mux.Handle("/cars", CarsRoute())

	return mux
}
