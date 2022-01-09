package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	. "user_information_service/data_loaders"
	. "user_information_service/models"
	. "user_information_service/utils"

	"github.com/gorilla/mux"
)

func HandleInterest(w http.ResponseWriter, r *http.Request) {

	data, err := LoadInterests()

	if err != nil {
		panic(err)
	}
	interestId := mux.Vars(r)["id"]
	if !IsEmpty(interestId) {
		interestIdInt, _ := strconv.Atoi(interestId)
		interestInList := findInInterestData(interestIdInt, data)

		if interestInList != nil {
			data = nil
			data = append(data, *interestInList)
		} else {
			data = nil
			w.WriteHeader(http.StatusNotFound)
		}

	}

	w.Header().Add("Content-type", "application-json")
	jsondata, _ := json.Marshal(data)
	w.Write(jsondata)
}

func findInInterestData(interestId int, data []Interest) *Interest {

	for _, item := range data {
		if item.ID == interestId {
			return &item

		}
	}
	return nil
}
