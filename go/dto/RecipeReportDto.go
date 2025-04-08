package dto

type RecipeReportDto struct {
	Moment string `json:"moment"`
	Type   string `json:"type"`
	Count  int    `json:"count"`
}

func NewMomentReport() []RecipeReportDto {
	var moments = []string{"Breakfast", "Lunch", "Snack", "Dinner"}
	var momentReport []RecipeReportDto
	for _, moment := range moments {
		momentReport = append(momentReport, RecipeReportDto{
			Type:   "",
			Moment: moment,
			Count:  0,
		})
	}
	return momentReport
}

func NewFoodReport() []RecipeReportDto {
	var types = []string{"Vegetable", "Fruit", "Cheese", "Dairy", "Meat"}
	var foodReport []RecipeReportDto

	for _, ftype := range types {
		foodReport = append(foodReport, RecipeReportDto{
			Type:   ftype,
			Moment: "",
			Count:  0,
		})
	}
	return foodReport
}
