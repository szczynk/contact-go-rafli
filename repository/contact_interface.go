package repository

import (
	"contact-go/model"
)

type ContactRepositorier interface {
	List() (result []model.Contact)
	Add(req model.ContactRequest) (contact model.Contact, err error)
	Update(id int, req model.ContactRequest) (err error)
	Delete(id int) (err error)
}