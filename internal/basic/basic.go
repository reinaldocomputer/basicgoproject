package basic

import (
	"github.com/reinaldocomputer/basicgoproject/internal/platform/mockDB"
	"sync"
)

var mu sync.Mutex

func NewBasic(r Request) *Basic {
	return &Basic{
		Id:   r.Id,
		Name: r.Name,
		Age:  r.Age,
	}
}

func (b Basic) Insert() error {
	mu.Lock()
	err := mockDB.InsertBasic(mockDB.BasicSchema{
		Id:   b.Id,
		Name: b.Name,
		Age:  b.Age,
	})
	mu.Unlock()
	return err
}

func (b Basic) DeleteByID() error {
	return mockDB.DeleteByID(b.Id)
}

func (b Basic) UpdateByID() (Basic, error) {
	data, err := mockDB.UpdateByID(b.Id, mockDB.BasicSchema{
		Id:   b.Id,
		Name: b.Name,
		Age:  b.Age,
	})
	return Basic{
		Id:   data.Id,
		Name: data.Name,
		Age:  data.Age,
	}, err
}

func (b Basic) GetByID() (Basic, error) {
	data, err := mockDB.GetByID(b.Id)
	return Basic{
		Id:   data.Id,
		Name: data.Name,
		Age:  data.Age,
	}, err
}

func GetAll() ([]Basic, error) {
	data, err := mockDB.GetAll()
	var dataBasic []Basic
	for _, d := range data {
		dataBasic = append(dataBasic, Basic{
			Id:   d.Id,
			Name: d.Name,
			Age:  d.Age,
		})
	}
	return dataBasic, err
}
