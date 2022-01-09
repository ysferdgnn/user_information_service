package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	. "user_information_service/data_loaders"
	. "user_information_service/models"
	"user_information_service/utils"

	"github.com/gorilla/mux"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {

	userId := mux.Vars(r)["id"]

	data, err := LoadUsers()
	if err != nil {
		panic(err)
	}

	if !utils.IsEmpty(userId) {

		userIdInt, _ := strconv.Atoi(userId)
		userInList := findUserInSlice(userIdInt, data)

		if userInList != nil {
			data = nil
			data = append(data, *userInList)
		} else {
			data = nil
			w.WriteHeader(http.StatusNotFound)
		}

	}

	jsondata, _ := json.Marshal(data)
	w.Header().Add("Content-type", "application-json")
	w.Write(jsondata)

}

func findUserInSlice(userId int, userlist []User) *User {

	for _, item := range userlist {
		if item.ID == userId {
			return &item
		}
	}
	return nil
}
