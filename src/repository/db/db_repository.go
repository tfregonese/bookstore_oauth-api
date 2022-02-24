package db

import (
	"github.com/gocql/gocql"
	"github.com/tfregonese/bookstore_oauth-api/src/clients/cassandra"
	"github.com/tfregonese/bookstore_oauth-api/src/domain/access_token"
	"github.com/tfregonese/bookstore_oauth-api/src/utils/errors"
)

func New() DdRepository {
	return &dbRepository{}
}

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_token(access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
	queryUpdateExpires     = "UPDATE access_token SET expires=? WHERE access_token=?"
)

type DdRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(s string) (*access_token.AccessToken, *errors.RestErr) {

	accessToken := access_token.GetNewAccessToken()
	if err := cassandra.GetSession().Query(queryGetAccessToken, s).Scan(
		&accessToken.AccessToken,
		&accessToken.UserId,
		&accessToken.ClientId,
		&accessToken.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError(err.Error())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &accessToken, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		token.AccessToken,
		token.UserId,
		token.ClientId,
		token.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryUpdateExpires,
		token.Expires,
		token.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError("Error while saving the Access Token")
	}

	return nil
}
