package usecase

import "collector/internal/domain"

type GitHubClient interface {
	GetRepositoryInfo(owner, name string) (domain.Repository, error)
}
