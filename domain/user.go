package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	ID               string    `json:"user_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	Name             string    `json:"name" valid:"notnull"`
	Email            string    `json:"status" valid:"notnull"`
	Error            string    `valid:"-"`
	CreatedAt        time.Time `json:"created_at" valid:"-"`
	UpdatedAt        time.Time `json:"updated_at" valid:"-"`
}

func NewUser(user *User) (*User, error) {

	user := User{
		User:            user,
	}

	user.prepare()

	err := user.Validate()

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (user *User) prepare() {
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (user *User) Validate() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}

	return nil
}
