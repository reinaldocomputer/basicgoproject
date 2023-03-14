package mockDB

import (
	"errors"
)

var (
	DataNotFoundByIDError = errors.New("There isn't any data with this ID.")
)

type BasicSchema struct {
	Id   int
	Name string
	Age  int
}

var database map[int]BasicSchema

func init() {
	database = make(map[int]BasicSchema)
}

func InsertBasic(data BasicSchema) error {
	database[data.Id] = data
	return nil
}

func GetAll() ([]BasicSchema, error) {
	var allData []BasicSchema

	for _, data := range database {
		allData = append(allData, data)
	}
	return allData, nil
}

func GetByID(id int) (BasicSchema, error) {
	got, ok := database[id]
	if ok {
		return got, nil
	}
	return BasicSchema{}, DataNotFoundByIDError
}

func UpdateByID(id int, updatedData BasicSchema) (BasicSchema, error) {
	_, ok := database[id]
	if ok {
		database[id] = updatedData
		return updatedData, nil
	}
	return BasicSchema{}, DataNotFoundByIDError
}

func DeleteByID(id int) error {
	_, ok := database[id]
	if ok {
		delete(database, id)
		return nil
	}
	return DataNotFoundByIDError
}
