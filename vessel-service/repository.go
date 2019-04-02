package main

import (
	"errors"

	"github.com/globalsign/mgo"

	pb "github.com/gpng/grpc-tutorial/vessel-service/proto/vessel"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

// Repository is an interface for repo
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

// VesselRepository is a temp repo
type VesselRepository struct {
	session *mgo.Session
}

// FindAvailable find available vessels
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessels []*pb.Vessel

	err := repo.collection().Find(nil).All(&vessels)

	if err != nil {
		return nil, err
	}

	for _, vessel := range vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that specification")
}

// Create new vessel
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

// Close closes db session
func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
