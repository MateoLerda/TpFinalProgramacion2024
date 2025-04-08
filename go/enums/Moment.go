package enums

type Moment int

const (
	InvalidMoment Moment = iota
	Breakfast
	Lunch
	Snack
	Dinner
)

func (m Moment) String() string {
	return []string{"InvalidMoment", "Breakfast", "Lunch", "Snack", "Dinner"}[m]
}

func ArrayString(moments []Moment) []string {
	var parsedResult []string
	for _, moment := range moments {
		parsedResult = append(parsedResult, []string{"InvalidMoment", "Breakfast", "Lunch", "Snack", "Dinner"}[moment])
	}
	return parsedResult
}

func GetMomentEnum(c string) Moment {
	switch c {
	case "Breakfast":
		return Breakfast
	case "Lunch":
		return Lunch
	case "Snack":
		return Snack
	case "Dinner":
		return Dinner
	default:
		return InvalidMoment
	}
}

func GetArrayMoments(listString []string) []Moment {
	var parsedResult []Moment
	for _, moment := range listString {
		switch moment {
		case "Breakfast":
			parsedResult = append(parsedResult, Breakfast)
		case "Lunch":
			parsedResult = append(parsedResult, Lunch)
		case "Snack":
			parsedResult = append(parsedResult, Snack)
		case "Dinner":
			parsedResult = append(parsedResult, Dinner)
		default:
			parsedResult = append(parsedResult, InvalidMoment)
		}
	}
	return parsedResult
}
