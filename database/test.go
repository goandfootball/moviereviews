package database

type DbLayer struct {
}

func (db *DbLayer) SelectTestOne(par int) string {
	return getValues(par)
}

func (db *DbLayer) SelectTestTwo(par int) string {
	return getValues(par)
}