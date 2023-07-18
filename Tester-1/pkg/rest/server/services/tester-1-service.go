package services

import (
	"github.com/sindhutrisha/Tester-1/tester-1/pkg/rest/server/daos"
	"github.com/sindhutrisha/Tester-1/tester-1/pkg/rest/server/models"
)

type Tester1Service struct {
	tester1Dao *daos.Tester1Dao
}

func NewTester1Service() (*Tester1Service, error) {
	tester1Dao, err := daos.NewTester1Dao()
	if err != nil {
		return nil, err
	}
	return &Tester1Service{
		tester1Dao: tester1Dao,
	}, nil
}

func (tester1Service *Tester1Service) CreateTester1(tester1 *models.Tester1) (*models.Tester1, error) {
	return tester1Service.tester1Dao.CreateTester1(tester1)
}

func (tester1Service *Tester1Service) UpdateTester1(id int64, tester1 *models.Tester1) (*models.Tester1, error) {
	return tester1Service.tester1Dao.UpdateTester1(id, tester1)
}

func (tester1Service *Tester1Service) DeleteTester1(id int64) error {
	return tester1Service.tester1Dao.DeleteTester1(id)
}

func (tester1Service *Tester1Service) ListTester1s() ([]*models.Tester1, error) {
	return tester1Service.tester1Dao.ListTester1s()
}

func (tester1Service *Tester1Service) GetTester1(id int64) (*models.Tester1, error) {
	return tester1Service.tester1Dao.GetTester1(id)
}
