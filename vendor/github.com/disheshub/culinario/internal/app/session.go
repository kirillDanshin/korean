package app

// Session - provides to userID
type Session struct {
	UserID int
}

// CheckAuth - method for getting userID
func (app *app) CheckAuth(token string) (*Session, error) {
	userID, err := app.session.GetID(token)
	if err != nil {
		return nil, err
	}

	return &Session{UserID: userID}, nil
}
