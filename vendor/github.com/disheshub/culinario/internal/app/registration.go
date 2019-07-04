package app

import "github.com/disheshub/culinario/internal/lib/hash"

// NewUser - data required for user create
type NewUser struct {
	Email    string
	Username string
	Pass     string
}

// Registration - method for create user
func (app *app) Registration(ctx Ctx, log Log, user NewUser) (token string, err error) {
	var ID int
	ID, err = app.db.CreateUser(ctx, log, user)
	if err != nil {
		return "", err
	}

	token = hash.RandToken()
	if err := app.session.Set(token, ID); err != nil {
		return "", err
	}

	return token, nil
}
