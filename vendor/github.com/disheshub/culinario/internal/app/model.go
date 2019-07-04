package app

type (
	User struct {
		ID       int
		Email    string
		Username string
		Avatar   string
	}

	MeasureWeight struct {
		ID   int
		Name string
	}

	Ingredient struct {
		ID   int
		Name string
	}

	Recipe struct {
		ID          int
		CookingTime int
		Description string
		Name        string
		Steps       []Step
		Ingredients []struct {
			Ingredient    Ingredient
			MeasureWeight MeasureWeight
			Count         int
		}
	}

	Step struct {
		Number int
		Text   string
	}
)
