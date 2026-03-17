package grpsserver

import (
	"collector/internal/usecase"
	pb "collector/proto"
	"context"
)

type Server struct {
	pb.UnimplementedCollectorServiceServer
	useCase usecase.GetRepositoryInfoUseCase
}

func NewServer(usecase usecase.GetRepositoryInfoUseCase) (s Server) {
	return Server{
		useCase: usecase,
	}
}
func (s Server) GetRepository(ctx context.Context, req *pb.RepositoryInfo) (*pb.Repository, error) {
	owner := req.Owner
	name := req.Name

	resp, err := s.useCase.GetRepositoryInfo(owner, name)
	if err != nil {
		return &pb.Repository{}, err
	}
	return &pb.Repository{
		Name:        resp.Name,
		Description: resp.Description,
		StarsCount:  int32(resp.Stars),
		Forks:       int32(resp.Forks),
		CreatedAt:   resp.CreatedAt,
	}, nil
}
