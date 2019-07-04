package app

func (app *app) GetRecipe(ctx Ctx, log Log, recipeID int) (*Recipe, error) {
	return app.db.GetRecipe(ctx, log, recipeID)
}
