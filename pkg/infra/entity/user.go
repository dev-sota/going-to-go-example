package entity

import (
	"time"

	"github.com/ispec-inc/going-to-go-example/pkg/domain/model"
)

type User struct {
	ID        int64      `gorm:"column:id; type:bigint(20) auto_increment; not null; primary_key"`
	Name      string     `gorm:"column:name"`
	Age       int        `gorm:"column:age"`
	CreatedAt *time.Time `gorm:"column:created_at; not null"`
	UpdatedAt *time.Time `gorm:"column:updated_at; not null"`
}

func NewUserFromModel(
	m model.User,
) User {
	return User{
		ID:   m.ID,
		Name: m.Name,
		Age:  m.Age,
	}
}

func (i User) ToModel() model.User {
	return model.User{
		ID:   i.ID,
		Name: i.Name,
		Age:  i.Age,
	}
}
