package main

import (
	"net/http"
	. "user_information_service/handlers"

	"github.com/gorilla/mux"
)

func main() {

	gorillaRoute := mux.NewRouter()

	gorillaRoute.HandleFunc("/users", HandleUser)
	gorillaRoute.HandleFunc("/users/", HandleUser)
	gorillaRoute.HandleFunc("/users/{id:[0-9]+}", HandleUser)

	gorillaRoute.HandleFunc("/interests", HandleInterest)
	gorillaRoute.HandleFunc("/interests/", HandleInterest)
	gorillaRoute.HandleFunc("/interests/{id:[0-9]+}", HandleInterest)

	gorillaRoute.HandleFunc("/usermapinterests", HandleUserMapInterests)
	gorillaRoute.HandleFunc("/usermapinterests/", HandleUserMapInterests)
	gorillaRoute.HandleFunc("/usermapinterests/{id:[0-9]+}", HandleUserMapInterests)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe("localhost:8081", nil)

}
