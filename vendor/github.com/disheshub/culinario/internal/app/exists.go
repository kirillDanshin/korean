package app

func (app *app) CheckEmail(ctx Ctx, log Log, email string) (exists bool, err error) {
	return app.db.ExistsEmail(ctx, log, email)
}
