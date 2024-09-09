package types

import (
	"time"

	"github.com/google/uuid"
)

type UserStore interface {
	GetUserByUsername(username string) (*User, error)
	GetUserByID(id uuid.UUID) (*User, error)
	GetUserByUUID(uuid uuid.UUID) (*User, error)
	CreateUser(User) error
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	CreateProduct(CreateProductPayload) error
	GetProductByID(id int) (*Product, error)
	DeleteProductByID(id int) error
}

type User struct {
	ID        uuid.UUID `json:"id"`
	UUID      uuid.UUID `json:"uuid"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       uuid.UUID `json:"image"`
	Price       int       `json:"price"`
	Permissions int       `json:"permission"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type RegisterUserPayload struct {
	UUID     uuid.UUID `json:"uuid" validate:"required,uuid"`
	Username string    `json:"username" validate:"required,max=63"`
	Password string    `json:"password" validate:"required,min=3,max=16"`
}

type LoginPayload struct {
	Username string `json:"username" validate:"required,max=63"`
	Password string `json:"password" validate:"required,min=3,max=16"`
}

type CreateProductPayload struct {
	Name        string    `json:"name" validate:"required,max=200"`
	Description string    `json:"description" validate:"required"`
	Image       uuid.UUID `json:"image" validate:"uuid"`
	Price       int       `json:"price" validate:"required"`
	Permissions int       `json:"permission"`
}
