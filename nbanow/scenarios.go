package scenarios

import (
	"ns-open-match/internal/util/testing"
	"ns-open-match/nbanow/scenarios/battle"
	"ns-open-match/pkg/matchfunction"
	"ns-open-match/pkg/pb"
	"sync"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	queryServiceAddress = "tasks.om-query:47103" // Address of the QueryService Endpoint.
	logger              = logrus.WithFields(logrus.Fields{
		"app": "nbanow",
	})
)

type GameScenarios interface {
	// Profiles lists all of the profiles that should run.
	Profiles() []*pb.MatchProfile

	// MatchFunction is the custom logic implementation of the match function.
	MatchFunction(p *pb.MatchProfile, poolTickets map[string][]*pb.Ticket) ([]*pb.Match, error)

	// Evaluate is the custom logic implementation of the evaluator.
	Evaluate(stream pb.Evaluator_EvaluateServer) error
}

var ActiveScenario = func() *Scenario {
	var gs GameScenarios = battle.Scenario()
	return &Scenario{
		Profiles:              gs.Profiles,
		MatchFunction:         queryPoolsWrapper(gs.MatchFunction),
		Evaluator:             gs.Evaluate,
		BackendDeletesTickets: false,
	}
}()

type fnMatch func(*pb.RunRequest, pb.MatchFunction_RunServer) error
type fnEvaluator func(pb.Evaluator_EvaluateServer) error

type Scenario struct {
	Profiles              func() []*pb.MatchProfile
	MatchFunction         fnMatch
	Evaluator             fnEvaluator
	BackendDeletesTickets bool
}

func (m fnMatch) Run(req *pb.RunRequest, server pb.MatchFunction_RunServer) error {
	return m(req, server)
}

func (e fnEvaluator) Evaluate(server pb.Evaluator_EvaluateServer) error {
	return e(server)
}

func getQueryServiceGRPCClient() pb.QueryServiceClient {
	conn, err := grpc.Dial(queryServiceAddress, testing.NewGRPCDialOptions(logger)...)
	if err != nil {
		logger.Fatalf("Failed to connect to Open Match, got %v", err)
	}
	return pb.NewQueryServiceClient(conn)
}

func queryPoolsWrapper(mmf func(req *pb.MatchProfile, pools map[string][]*pb.Ticket) ([]*pb.Match, error)) fnMatch {
	var q pb.QueryServiceClient
	var startQ sync.Once

	return func(req *pb.RunRequest, stream pb.MatchFunction_RunServer) error {
		startQ.Do(func() {
			q = getQueryServiceGRPCClient()
		})

		//logger.Infof("START QueryPools..")
		poolTickets, err := matchfunction.QueryPools(stream.Context(), q, req.GetProfile().GetPools())
		if err != nil {
			return err
		}

		//for pool, tickets := range poolTickets {
		//	logger.Infof("Pool(%v) Tickets(%v)\n", pool, tickets)
		//}

		// 매칭Pool에서 유저 매칭
		proposals, err := mmf(req.GetProfile(), poolTickets)
		if err != nil {
			return err
		}

		logger.WithFields(logrus.Fields{
			"proposals": proposals,
		}).Trace("proposals returned by match function")

		// 매칭된 정보를 BackEnd서버로 보낸다.
		for _, proposal := range proposals {
			if err := stream.Send(&pb.RunResponse{Proposal: proposal}); err != nil {
				return err
			}
		}

		return nil
	}
}
