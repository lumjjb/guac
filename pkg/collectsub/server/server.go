package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/guacsec/guac/pkg/collectsub/collectsub"
	"github.com/guacsec/guac/pkg/collectsub/server/db/simpledb"
	db "github.com/guacsec/guac/pkg/collectsub/server/db/types"
	"github.com/guacsec/guac/pkg/logging"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedColectSubscriberServiceServer

	// Db points to the backend DB, public for mocking testing purposes.
	Db   db.CollectSubscriberDb
	port int
}

func NewServer(port int) (*server, error) {
	db, err := simpledb.NewSimpleDb()
	if err != nil {
		return nil, err
	}

	return &server{
		Db:   db,
		port: port,
	}, nil
}

func (s *server) AddCollectEntry(ctx context.Context, in *pb.AddCollectEntriesRequest) (*pb.AddCollectEntriesResponse, error) {
	err := s.Db.AddCollectEntries(ctx, in.Entries)
	if err != nil {
		return nil, fmt.Errorf("failed to add entry to db: %w", err)
	}

	return &pb.AddCollectEntriesResponse{
		Success: true,
	}, nil
}

func (s *server) GetCollectEntries(ctx context.Context, in *pb.GetCollectEntriesRequest) (*pb.GetCollectEntriesResponse, error) {
	ret, err := s.Db.GetCollectEntries(ctx, in.Filters)
	if err != nil {
		return nil, fmt.Errorf("failed to get collect entries from db: %w", err)
	}
	return &pb.GetCollectEntriesResponse{
		Entries: ret,
	}, nil
}

func (s *server) Serve(ctx context.Context) error {
	logger := logging.FromContext(ctx)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	gs := grpc.NewServer()
	pb.RegisterColectSubscriberServiceServer(gs, s)
	logger.Infof("server listening at %v", lis.Addr())
	if err := gs.Serve(lis); err != nil {
		return err
	}

	return nil
}
