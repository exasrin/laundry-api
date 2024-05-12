package manager

import (
	"go-api-enigma/usecase"
)

type UseCaseManager interface {
	UomUseCase() usecase.UomUseCase
	// ProductUseCase() usecase.ProductUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

// UomUseCase implements UseCaseManager.
func (u *useCaseManager) UomUseCase() usecase.UomUseCase {
	return usecase.NewUomUseCase(u.repoManager.UomRepository())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
