// Each file must belong to a package to do this at the beginning of the file the name of the package is declared.
package main

import (
	"net/http"
	"raw-service/routes"
)

// import modules. _ means that it's not necessary to use the module, in this case we only need to load data in memory before starting the server.

// main function is the first to be executed when the program run.
func main() {
	routes := routes.LoadRoutes()

	http.ListenAndServe(":8080", routes)
}
