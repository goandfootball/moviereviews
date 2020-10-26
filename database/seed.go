package database

func getValues (par int) string{
	var result string

	if par == 1 {
		result = "string value one"
	} else if par == 2 {
		result = "string value two"
	} else {
		result = "not parameter value identified"
	}

	return result
}