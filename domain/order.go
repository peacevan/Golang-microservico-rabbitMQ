package domain

import (
	"github.com/asaskevich/govalidator"
	"time"
)

id, 
user_id, 
pair, 
amount, 
direction, 
type (market, limit), 
created_at, 
updated_at

type Order struct {
	ID                     string     `json:"id_order" valid:"uuid" gorm:"type:uuid;primary_key"`
	user_id                string     `json:"user_id" valid:"notnull" gorm:"type:varchar(255)"`
	pair                   string     `json:"pair" valid:"notnull" gorm:"type:varchar(255)"`
	amount                 string     `json:"amount" valid:"notnull" gorm:"type:varchar(255)"`
	type_order             string     `json:"type" valid:"notnull" gorm:"type:varchar(255)"`
	created_at  time.Time  `json:"-" valid:"-"`
	updated_at  time.Time  `json:"-" valid:"-"`
	user       []*User     `json:"-" valid:"-" gorm:"ForeignKey:OrderID"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewOder() *Order{
	return &Order{}
}

func (order *Order) Validate() error {

	_, err := govalidator.ValidateStruct(order)

	if err != nil {
		return err
	}

	return nil
}
