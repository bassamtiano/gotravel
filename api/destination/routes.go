package catalogs

import (
	destinationController "github.com/bassamtiano/gotravel/api/destination/controllers"
	"github.com/gorilla/mux"
)

func AddDestinationRouterHandler(r *mux.Router) {
	url_prefix := "/destination/"

	r.HandleFunc(url_prefix+"hello", destinationController.HelloWorldHandler).Methods("GET")

}
