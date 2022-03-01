package rest

import (
	"encoding/json"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/tfregonese/bookstore_oauth-api/src/domain/users"
	"github.com/tfregonese/bookstore_oauth-api/src/utils/errors"
	"time"
)

var (
	userRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type userRepository struct {
}

func NewRepository() RestUsersRepository {
	return &userRepository{}
}

func (ur *userRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {

	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := userRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid Rest Client Response when trying to login user")
	}

	// error situation
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface whe trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshall users response")
	}

	return &user, nil
}
