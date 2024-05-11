package usecase

import (
	"fmt"
	"go-api-enigma/model"
	"go-api-enigma/repository"
)

type UomUseCase interface {
	CreateNew(payload model.Uom) error
	FindById(id string) (model.Uom, error)
	FindAll() ([]model.Uom, error)
	Update(payload model.Uom) error
	Delete(id string) error
}

type uomUseCase struct {
	repo repository.UomRepository
}

// CreateNew implements UomUseCase.
func (u *uomUseCase) CreateNew(payload model.Uom) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Type == "" {
		return fmt.Errorf("type is required")
	}

	err := u.repo.Save(payload)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements UomUseCase.
func (u *uomUseCase) Delete(id string) error {
	uom, err := u.repo.FindById(id)
	if err != nil {
		return fmt.Errorf("uom not found")
	}

	err = u.repo.DeleteById(uom.Id)
	if err != nil {
		return fmt.Errorf("failed to delete uom: %v", err)
	}
	return nil
}

// FindAll implements UomUseCase.
func (u *uomUseCase) FindAll() ([]model.Uom, error) {
	uoms, err := u.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all: %v", err)
	}
	return uoms, nil
}

// FindById implements UomUseCase.
func (u *uomUseCase) FindById(id string) (model.Uom, error) {
	uom, err := u.repo.FindById(id)
	if err != nil {
		return model.Uom{}, fmt.Errorf("uom not found")
	}
	return uom, nil
}

// Update implements UomUseCase.
func (u *uomUseCase) Update(payload model.Uom) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Type == "" {
		return fmt.Errorf("type is required")
	}

	_, err := u.repo.FindById(payload.Id)
	if err != nil {
		return err
	}

	err = u.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update uom: %v", err)
	}
	return nil
}

func NewUomUseCase(repo repository.UomRepository) UomUseCase {
	return &uomUseCase{repo: repo}
}
