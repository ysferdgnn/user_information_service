package viewmodels

import (
	. "user_information_service/models"
)

type UserMapInterestViewModel struct {
	User      User
	Interests []Interest
}
