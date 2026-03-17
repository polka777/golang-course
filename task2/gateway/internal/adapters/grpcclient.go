package adapters

import (
	"context"
	"errors"
	"fmt"
	"gateway/internal/domain"
	pb "gateway/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcClient struct {
	conn   *grpc.ClientConn
	client pb.CollectorServiceClient
}

func NewGrpcClient(address string) (GrpcClient, error) {
	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second))
	if err != nil {
		return GrpcClient{}, fmt.Errorf("failed to connect to collector service: %w", err)
	}
	client := pb.NewCollectorServiceClient(conn)
	return GrpcClient{conn: conn, client: client}, nil
}
func (c GrpcClient) GetRepositoryInfo(ctx context.Context, owner, name string) (domain.Repository, error) {
	req := &pb.RepositoryInfo{
		Name:  name,
		Owner: owner,
	}
	resp, err := c.client.GetRepository(ctx, req)
	if err != nil {
		return domain.Repository{}, GrpcErr(err)
	}
	return domain.Repository{
		Name:        resp.Name,
		Description: resp.Description,
		Stars:       int(resp.StarsCount),
		Forks:       int(resp.Forks),
		CreatedAt:   resp.CreatedAt,
	}, nil

}
func (c GrpcClient) Close() error {
	return c.conn.Close()
}
func GrpcErr(err error) error {
	if err == nil {
		return nil
	}
	status, ok := status.FromError(err)
	if !ok {
		return fmt.Errorf("unknown grpc error: %w", err)
	}
	switch status.Code() {
	case codes.NotFound:
		return errors.New("repository not found")
	case codes.InvalidArgument:
		return errors.New("invalid argument")
	case codes.Unavailable:
		return errors.New("grpc server unavailable")
	default:
		return fmt.Errorf("grps error: %w", err)
	}

}
