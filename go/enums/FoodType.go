package enums

type FoodType int

const (
	InvalidFoodType FoodType = iota
	Vegetable
	Fruit
	Cheese
	Dairy
	Meat
)

func (f FoodType) String() string {
	return []string{"InvalidFoodType","Vegetable", "Fruit", "Cheese", "Dairy", "Meat"}[f]
}

func GetTypeEnum(c string) FoodType {
	switch c {
	case "Vegetable":
		return Vegetable
	case "Fruit":
		return Fruit
	case "Cheese":
		return Cheese
	case "Dairy":
		return Dairy
	case "Meat":
		return Meat
	default:
		return InvalidFoodType
	}
}
