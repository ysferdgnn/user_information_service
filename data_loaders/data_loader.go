package dataloaders

import (
	"encoding/json"
	"log"
	. "user_information_service/models"
	. "user_information_service/utils"
)

const userfilepath string = "../json/user.json"
const InterestFilePath string = "../json/interest.json"
const userMapInterestFilePath string = "../json/user_map_interest.json"

func LoadUsers() ([]User, error) {

	log.Printf("Users loading from external file. Path is -> %s", userfilepath)
	var users []User

	data, dataReadError := ReadFile(userfilepath)

	if dataReadError != nil {
		log.Fatal(dataReadError)
		return nil, dataReadError
	}

	err := json.Unmarshal(data, &users)

	if err != nil {
		log.Fatalf("An error has occurred. Data can not unmarshal")
		log.Fatal(err)
		return nil, err
	}

	return users, nil

}

func LoadInterests() ([]Interest, error) {

	log.Printf("Interest loading from external file. Path is -> %s", userfilepath)
	var interests []Interest

	data, dataReadError := ReadFile(InterestFilePath)

	if dataReadError != nil {
		log.Fatal(dataReadError)
		return nil, dataReadError
	}

	err := json.Unmarshal(data, &interests)

	if err != nil {
		log.Fatalf("An error has occurred. Data can not unmarshal")
		log.Fatal(err)
		return nil, err
	}

	return interests, nil

}

func LoadUserMapInterest() ([]UserMapInterest, error) {

	log.Printf("User Map Interest loading from external file. Path is -> %s", userfilepath)
	var userMapInterests []UserMapInterest

	data, dataReadError := ReadFile(userMapInterestFilePath)

	if dataReadError != nil {
		log.Fatal(dataReadError)
		return nil, dataReadError
	}

	err := json.Unmarshal(data, &userMapInterests)

	if err != nil {
		log.Fatalf("An error has occurred. Data can not unmarshal")
		log.Fatal(err)
		return nil, err
	}

	return userMapInterests, nil

}
