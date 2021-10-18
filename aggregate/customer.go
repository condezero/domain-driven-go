package aggregate

import (
	"errors"

	"github.com/condezero/domain-driven-go/entity"
	"github.com/condezero/domain-driven-go/valueobject"
	"github.com/google/uuid"
)

var (
	// ErrInvalidPerson is returned when the person is not valid in the NewCustome factory
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

type Customer struct {
	person *entity.Person

	products []*entity.Item

	transactions []*valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]*valueobject.Transaction, 0),
	}, nil
}
