package models

import (
	"context"
	"errors"

	"github.com/go-fuego/fuego"
)

type Users struct {
	ID        string `json:"id" validate:"required" example:"user-123456"`
	Name      string `json:"name" validate:"required" example:"Napoleon"`
	Age       int    `json:"age" example:"2" description:"Age of the user, in years"`
	IsAdopted bool   `json:"is_adopted" description:"Is the user adopted"`
}

type UsersCreate struct {
	Name      string `json:"name" validate:"required,min=1,max=100" example:"Napoleon"`
	Age       int    `json:"age" validate:"min=0,max=100" example:"2" description:"Age of the user, in years"`
	IsAdopted bool   `json:"is_adopted" description:"Is the user adopted"`
}

type UsersUpdate struct {
	Name      string `json:"name,omitempty" validate:"min=1,max=100" example:"Napoleon" description:"Name of the user"`
	Age       int    `json:"age,omitempty" validate:"max=100" example:"2"`
	IsAdopted *bool  `json:"is_adopted,omitempty" description:"Is the user adopted"`
}

var _ fuego.InTransformer = &Users{}

func (*Users) InTransform(context.Context) error {
	return errors.New("users must only be used as output")
}
