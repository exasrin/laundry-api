package repository

import (
	"database/sql"
	"go-api-enigma/model"
)

type ProductRepository interface {
	Save(product model.Product) error
	FindById(id string) (model.Product, error)
	FindAll() ([]model.Product, error)
	Update(product model.Product) error
	DeleteById(id string) error
}

type productRepository struct {
	db *sql.DB
}

// DeleteById implements ProductRepository.
func (p *productRepository) DeleteById(id string) error {
	panic("unimplemented")
}

// FindAll implements ProductRepository.
func (p *productRepository) FindAll() ([]model.Product, error) {
	panic("unimplemented")
}

// FindById implements ProductRepository.
func (p *productRepository) FindById(id string) (model.Product, error) {
	panic("unimplemented")
}

// Save implements ProductRepository.
func (p *productRepository) Save(product model.Product) error {
	_, err := p.db.Exec("INSERT INTO m_product VALUES($1, $2, $3, $4)", product.Id, product.Name, product.Price, product.Uom.Id)
	if err != nil {
		return err
	}
	return nil
}

// Update implements ProductRepository.
func (p *productRepository) Update(product model.Product) error {
	panic("unimplemented")
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
