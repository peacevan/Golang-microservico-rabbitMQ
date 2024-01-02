package repositories

import (
	"encoder/domain"
	"fmt"
	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	Insert(order *domain.Order) (*domain.Order, error)
	Find(id string) (*domain.Order, error)
	Update(order *domain.Order) (*domain.Order, error)
	//FindAll() ([]domain.Order, error)
}

type OrderRepositoryDb struct {
	Db *gorm.DB
}

func (repo OrderRepositoryDb) Insert(order *domain.Order) (*domain.Order, error) {

	err := repo.Db.Create(order).Error

	if err != nil {
		return nil, err
	}

	return order, nil

}

func (repo OrderRepositoryDb) Find(id string) (*domain.Order, error) {

	var order domain.Order
	repo.Db.Preload("Video").First(&order, "id = ?", id)

	if order.ID == "" {
		return nil, fmt.Errorf("order does not exist")
	}

	return &order, nil
}

func (repo OrderRepositoryDb) Update(order *domain.Order) (*domain.order, error) {
	err := repo.Db.Save(&order).Error

	if err != nil {
		return nil, err
	}

	return order, nil
}
