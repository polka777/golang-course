package usecase

import (
	"context"
	"errors"
	"gateway/internal/domain"
)

type GatewayUsecase struct {
	client CollectorClient
}

func NewGatewayUsecase(c CollectorClient) GatewayUsecase {
	return GatewayUsecase{client: c}
}
func (c GatewayUsecase) GetRepositoryInfo(cnx context.Context, name, owner string) (domain.Repository, error) {
	if name == "" || owner == "" {
		return domain.Repository{}, errors.New("the fields name and owner cannot be empty")
	}
	repo, err := c.client.GetRepositoryInfo(cnx, name, owner)
	if err != nil {
		return domain.Repository{}, err
	}
	return repo, nil
}
