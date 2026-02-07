package services

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repositories"
	"strings"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(data *models.Product) error {
	err := s.repo.Create(data)
	if err != nil && strings.Contains(err.Error(), "violates foreign key constraint") {
		return errors.New("invalid category_id: category not found")
	}

	return err
}

func (s *ProductService) GetAll(name string) ([]models.Product, error) {
	return s.repo.GetAll(name)
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
