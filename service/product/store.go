package product

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/afallenhope/go-vendor/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}
	return products, nil
}

func (s *Store) GetProductByID(id int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id = $1 LIMIT 1;", id)
	if err != nil {
		return nil, err
	}

	product := new(types.Product)
	for rows.Next() {
		product, err = scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	if product.ID == 0 {
		return nil, fmt.Errorf("product not found")
	}

	return product, nil
}

func (s *Store) DeleteProductByID(id int) error {
	_, err := s.db.Exec("DELETE FROM products WHERE id = $1;", id)
	if err != nil {
		log.Fatalf("error with sql %v", err)
		return err
	}

	return nil
}

func (s *Store) CreateProduct(product types.CreateProductPayload) error {
	_, err := s.db.Exec(
		"INSERT INTO products (name, description, image, price, permissions) VALUES ($1, $2, $3, $4, $5);",
		product.Name,
		product.Description,
		product.Image,
		product.Price,
		product.Permissions,
	)

	if err != nil {
		log.Fatalf("error with sql %v", err)
		return err
	}

	return nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Permissions,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
