package usecase

import (
	"context"
	"gateway/internal/domain"
)

type CollectorClient interface {
	GetRepositoryInfo(cnx context.Context, owner, name string) (domain.Repository, error)
}
