package dto

import "gateway/internal/domain"

type RepositoryResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
	Forks       int    `json:"forks"`
	CreatedAt   string `json:"created_at"`
}

func FromDomain(r domain.Repository) RepositoryResponse {
	return RepositoryResponse{
		Name:        r.Name,
		Description: r.Description,
		Stars:       r.Stars,
		Forks:       r.Forks,
		CreatedAt:   r.CreatedAt,
	}
}
