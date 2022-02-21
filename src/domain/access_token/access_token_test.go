package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	//utilizando go puro
	if expirationTime != 24 {
		t.Error("Expiration time should be 24 hours")
	}
	//utilizando la libreria assert
	assert.EqualValues(t, 24, expirationTime, "Expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	// Utilizando GO puro
	if at.isExpired() {
		t.Error("New Access Token should not be Expired.")
	}
	if at.AccessToken != "" {
		t.Error("New Access Token AccessToken should not be Set.")
	}
	if at.UserId != 0 {
		t.Error("New Access Token UserId should not be Set.")
	}

	// utilizando la libreria
	assert.False(t, at.isExpired(), "New Access Token should not be Expired.")
	assert.EqualValues(t, "", at.AccessToken, "New Access Token AccessToken should not be Set.")
	assert.EqualValues(t, 0, at.UserId, "New Access Token UserId should not be Set.")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	//utilizando go puro
	if !at.isExpired() {
		t.Error("empty access token should be expired by default")
	}
	//utilizando la libreria
	assert.True(t, at.isExpired(), "empty access token should be expired by default")

	//utilizando go puro
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.isExpired() {
		t.Error("Access token should not be expired")
	}
	//utilizando la libreria
	assert.False(t, at.isExpired(), "Access token should not be expired")
}
