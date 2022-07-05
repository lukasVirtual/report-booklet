package token

import (
	"encoding/base64"

	"www.github.com/backend/dbmodels"
)

func CreateBearer() *base64.Encoding {
	var token *base64.Encoding
	var user dbmodels.User
	token = base64.NewEncoding(user.Name + user.Name)
	return token
}
