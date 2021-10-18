package customer

import (
	"errors"

	"github.com/condezero/domain-driven-go/aggregate"
	"github.com/google/uuid"
)

var (
	// ErrCustomerNotFound is returned when a customer is not found
	ErrCustomerNotFound = errors.New("customer not found")
	// ErrFailedToCreate is returned when a customer is not created
	ErrFailedToCreate = errors.New("failed to create customer")
	// ErrFailedToUpdate is returned when a customer is not updated
	ErrFailedToUpdate = errors.New("failed to update customer")
)

// CustomerRepository is the interface for the customer repository
type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
