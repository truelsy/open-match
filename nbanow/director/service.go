package director

import (
	"context"
	"ns-open-match/internal/config"
	"ns-open-match/internal/statestore"
	"ns-open-match/pkg/pb"

	"github.com/golang/protobuf/ptypes/empty"
)

type service struct {
	cfg   config.View
	store statestore.Service
}

func (s *service) RegisterArena(c context.Context, syn *pb.SynRegisterArena) (*empty.Empty, error) {
	logger.Infof("Register!! syn(%+v)", *syn)
	return &empty.Empty{}, nil
}

func (s *service) HealthCheckArena(c context.Context, syn *pb.SynHealthCheckArena) (*empty.Empty, error) {
	logger.Infof("HealthCheck!! syn(%+v)", *syn)
	return &empty.Empty{}, nil
}

func (s *service) TerminateArena(c context.Context, syn *pb.SynTerminateArena) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *service) AvailableArena(c context.Context, syn *pb.SynAvailableArena) (*pb.AckAvailableArena, error) {
	return &pb.AckAvailableArena{
		Address: "",
	}, nil
}
