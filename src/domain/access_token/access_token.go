package access_token

import (
	"github.com/tfregonese/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grandTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrandType string `json:"grand_type"`
	Scope     string `json:"scope"`

	//Used in password grand type
	Username string `json:"username"`
	Password string `json:"password"`

	//Used in client_credentials grand type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrandType {
	case grandTypePassword:
		break
	case grandTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid Grand Type")
	}
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

//Different client_id, so you can do different things depending were the data is requested
// WEB frontend -> ClientId: 123
// Android API -> ClientId: 234

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if len(at.AccessToken) == 0 {
		return errors.NewBadRequestError("invalid AccessToken Id.")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid User Id.")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid Client Id.")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid Expires.")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) isExpired() bool {
	// return time.Now().UTC().Unix() > at.Expires
	now := time.Now().UTC()
	timeAccessTokenExpires := time.Unix(at.Expires, 0)
	return now.After(timeAccessTokenExpires)
}
