package main

import (
	"context"

	"github.com/globalsign/mgo"

	pb "github.com/gpng/grpc-tutorial/vessel-service/proto/vessel"
)

// our grpc service handler
type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	if err := repo.Create(req); err != nil {
		return err
	}

	res.Vessel = req
	res.Created = true
	return nil
}
