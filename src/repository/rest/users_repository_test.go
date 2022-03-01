package rest

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test cases ...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutAPI(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@gmail.com,"password":"password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := userRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	fmt.Println(user)
	fmt.Println(err)
}

func TestLoginInvalidErrorInterface(t *testing.T) {

}

func TestLoginInvalidLoginCredentials(t *testing.T) {

}

func TestLoginInvalidUserJsonResponse(t *testing.T) {

}

func TestLoginNoError(t *testing.T) {

}
