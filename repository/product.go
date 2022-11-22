package repository

import (
	"github.com/Taras-Rm/aws_rds/models"
	"github.com/go-pg/pg/v10"
)

type ProductModel struct {
	tableName struct{} `pg:"products"`
	ID        uint64
	Name      string `pg:"name"`
	Price     uint64 `pg:"price"`
}

type productRepository struct {
	db *pg.DB
}

type ProductInterface interface {
	Create(product ProductModel) error
	GetAll() ([]*models.Product, error)
}

func NewProductRepository(db *pg.DB) ProductInterface {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product ProductModel) error {
	_, err := r.db.Model(&product).Insert()
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) GetAll() ([]*models.Product, error) {
	var prodModels []ProductModel
	err := r.db.Model(&prodModels).Select()
	if err != nil {
		return nil, err
	}

	var products []*models.Product
	for _, m := range prodModels {
		products = append(products, &models.Product{
			Id:    m.ID,
			Name:  m.Name,
			Price: m.Price,
		})
	}

	return products, nil
}
