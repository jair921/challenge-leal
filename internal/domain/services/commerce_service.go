package services

import (
	"errors"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
)

// CommerceService define los métodos relacionados con la lógica de negocio de comercios.
type CommerceService interface {
	CreateCommerce(commerce *entities.Commerce) error
	GetCommerceByID(id string) (*entities.Commerce, error)
	ListCommerces() ([]*entities.Commerce, error)
}

// commerceService es una implementación concreta de CommerceService
type commerceService struct {
	commerceRepo repositories.CommerceRepository
}

// NewCommerceService crea un nuevo servicio de comercios
func NewCommerceService(commerceRepo repositories.CommerceRepository) CommerceService {
	return &commerceService{
		commerceRepo: commerceRepo,
	}
}

// CreateCommerce crea un nuevo comercio
func (s *commerceService) CreateCommerce(commerce *entities.Commerce) error {
	if commerce.Name == "" {
		return errors.New("name is required")
	}
	return s.commerceRepo.Create(commerce)
}

// GetCommerceByID obtiene un comercio por su ID
func (s *commerceService) GetCommerceByID(id string) (*entities.Commerce, error) {
	commerce, err := s.commerceRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return commerce, nil
}

// ListCommerces obtiene todos los comercios
func (s *commerceService) ListCommerces() ([]*entities.Commerce, error) {
	return s.commerceRepo.List()
}
