package app


func (app *app) RemoveIngredient(ctx Ctx, log Log, ID int) error {
	return app.db.RemoveIngredient(ctx, log, ID)
}
