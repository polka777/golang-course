package usecase

import (
	"collector/internal/domain"
	"errors"
)

type GetRepositoryInfoUseCase struct {
	gitHubClient GitHubClient
}

func NewGetRepositoryInfoUseCase(client GitHubClient) GetRepositoryInfoUseCase {
	return GetRepositoryInfoUseCase{gitHubClient: client}
}
func (r GetRepositoryInfoUseCase) GetRepositoryInfo(owner, name string) (domain.Repository, error) {
	if owner == "" || name == "" {
		return domain.Repository{}, errors.New("The fields name and owner cannot be empty!")
	}
	repo, error := r.gitHubClient.GetRepositoryInfo(owner, name)
	if error != nil {
		return domain.Repository{}, error
	}
	return repo, error
}
