package database

import "github.com/zorasantos/my-health/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
	Update(user *entity.User) error
}

type ExaminationInterface interface {
	Create(examination *entity.Examination) error
}
