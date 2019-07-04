package app

import (
	"github.com/disheshub/culinario/internal/lib/hash"
)

// LoginInfo - data for authorization user
type LoginInfo struct {
	Email, Password string
}

// Login - method for authorization user
func (app *app) Login(ctx Ctx, log Log, loginInfo LoginInfo) (token string, err error) {
	ID, passHash, err := app.db.GetPassAndIDByEmail(ctx, log, loginInfo.Email)
	if err != nil {
		return "", err
	}

	if !hash.CheckPasswordHash(loginInfo.Password, passHash) {
		return "", ErrPassNotMatch
	}

	token = hash.RandToken()
	if err = app.session.Set(token, ID); err != nil {
		return "", err
	}

	return token, nil
}
