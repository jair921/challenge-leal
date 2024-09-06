package services

import (
	"errors"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
)

// BranchService define los métodos relacionados con la lógica de negocio de las sucursales.
type BranchService interface {
	CreateBranch(branch *entities.Branch) error
	GetBranchByID(id string) (*entities.Branch, error)
	ListBranches() ([]*entities.Branch, error)
}

// branchService es una implementación concreta de BranchService
type branchService struct {
	branchRepo repositories.BranchRepository
}

// NewBranchService crea un nuevo servicio de sucursales
func NewBranchService(branchRepo repositories.BranchRepository) BranchService {
	return &branchService{
		branchRepo: branchRepo,
	}
}

// CreateBranch crea una nueva sucursal
func (s *branchService) CreateBranch(branch *entities.Branch) error {
	if branch.Name == "" || branch.CommerceID == "" {
		return errors.New("name and commerce ID are required")
	}
	return s.branchRepo.Create(branch)
}

// GetBranchByID obtiene una sucursal por su ID
func (s *branchService) GetBranchByID(id string) (*entities.Branch, error) {
	branch, err := s.branchRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return branch, nil
}

// ListBranches obtiene todas las sucursales
func (s *branchService) ListBranches() ([]*entities.Branch, error) {
	return s.branchRepo.List()
}
