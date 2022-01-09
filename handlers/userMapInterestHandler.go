package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	. "user_information_service/data_loaders"
	. "user_information_service/models"
	. "user_information_service/viewModels"

	"github.com/gorilla/mux"
)

func HandleUserMapInterests(w http.ResponseWriter, r *http.Request) {
	data, err := LoadUserMapInterest()
	userdata, err := LoadUsers()
	interestdata, err := LoadInterests()
	urlParameters := mux.Vars(r)
	userId := urlParameters["id"]
	userMapInterestViewModels := make([]UserMapInterestViewModel, 0)

	if err != nil {
		panic(err)
	}
	filteredUserMapInterest := make([]UserMapInterest, len(data))
	copy(filteredUserMapInterest, data)

	if len(userId) > 0 {

		filteredUserMapInterest = nil
		log.Printf("userId filtered ->%s", userId)
		userIdint, _ := strconv.Atoi(userId)
		for i := range data {
			if data[i].UserID == userIdint {
				filteredUserMapInterest = append(filteredUserMapInterest, data[i])
			}
		}

		for _, user := range userdata {
			if user.ID == userIdint {
				tmpUser := user
				userdata = nil
				userdata = append(userdata, tmpUser)
			}
		}
	}

	for _, user := range userdata {

		interestSlice := make([]Interest, 0)

		for _, usermapInterest := range filteredUserMapInterest {
			if user.ID == usermapInterest.UserID {
				for _, interest := range interestdata {
					if interest.ID == usermapInterest.InterestID {
						interestSlice = append(interestSlice, interest)
					}

				}
			}

		}
		tmpUserMapInterestViewModel := UserMapInterestViewModel{User: user, Interests: interestSlice}
		userMapInterestViewModels = append(userMapInterestViewModels, tmpUserMapInterestViewModel)

	}

	log.Printf("Requested userId ->%s", userId)
	jsondata, _ := json.Marshal(userMapInterestViewModels)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsondata)
}
