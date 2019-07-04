package app

type NewRecipe struct {
	UserID      int
	CookingTime int
	Description string
	Name        string
	Ingredients []RecipeIngredient
	Steps       []Step
}

type RecipeIngredient struct {
	ID, Count int
}

func (app *app) CreateRecipe(ctx Ctx, log Log, recipeInfo NewRecipe) (*Recipe, error) {
	return app.db.CreateRecipe(ctx, log, recipeInfo)
}
