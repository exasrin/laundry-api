package manager

import "go-api-enigma/repository"

type RepoManager interface {
	UomRepository() repository.UomRepository
	ProductRepository() repository.ProductRepository
}

type repoManager struct {
	infraManager InfraManager
}

// ProductRepository implements RepoManager.
func (r *repoManager) ProductRepository() repository.ProductRepository {
	return repository.NewProductRepository(r.infraManager.Conn())
}

// UomRepository implements RepoManager.
func (r *repoManager) UomRepository() repository.UomRepository {
	return repository.NewUomRepository(r.infraManager.Conn())
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
