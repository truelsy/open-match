package director

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"ns-open-match/internal/appmain"
	"ns-open-match/internal/config"
	"ns-open-match/internal/rpc"
	"ns-open-match/internal/telemetry"
	scenarios "ns-open-match/nbanow"
	"ns-open-match/pkg/pb"
	"sync"
	"time"

	"go.opencensus.io/trace"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

var (
	logger = logrus.WithFields(logrus.Fields{
		"app":       "openmatch",
		"component": "nbanow.director",
	})

	mIterations          = telemetry.Counter("nbanow_backend_iterations", "fetch match iterations")
	mFetchMatchCalls     = telemetry.Counter("nbanow_backend_fetch_match_calls", "fetch match calls")
	mFetchMatchSuccesses = telemetry.Counter("nbanow_backend_fetch_match_successes", "fetch match successes")
	mFetchMatchErrors    = telemetry.Counter("nbanow_backend_fetch_match_errors", "fetch match errors")
	mMatchesReturned     = telemetry.Counter("nbanow_backend_matches_returned", "matches returned")
	mSumTicketsReturned  = telemetry.Counter("nbanow_backend_sum_tickets_returned", "tickets in matches returned")
	mMatchesAssigned     = telemetry.Counter("nbanow_backend_matches_assigned", "matches assigned")
	mMatchAssignsFailed  = telemetry.Counter("nbanow_backend_match_assigns_failed", "match assigns failed")
	//mTicketsDeleted      = telemetry.Counter("nbanow_backend_tickets_deleted", "tickets deleted")
	//mTicketDeletesFailed = telemetry.Counter("nbanow_backend_ticket_deletes_failed", "ticket deletes failed")
)

func BindService(p *appmain.Params, b *appmain.Bindings) error {
	profiles := scenarios.ActiveScenario.Profiles()
	go doRun(p.Config(), profiles)
	return nil
}

func doRun(cfg config.View, profiles []*pb.MatchProfile) {
	conn, err := rpc.GRPCClientFromConfig(cfg, "api.backend")
	if err != nil {
		log.Fatalf("Failed to connect to Open Match Backend, got %s", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)

	be := pb.NewBackendServiceClient(conn)

	matchesForAssignment := make(chan *pb.Match, 30000)
	for i := 0; i < 50; i++ {
		go doAssignments(be, matchesForAssignment)
	}

	// Don't go faster than this, as it likely means that FetchMatches is throwing
	// errors, and will continue doing so if queried very quickly.
	for range time.Tick(time.Millisecond * 250) {
		// Keep pulling matches from Open Match backend
		var wg sync.WaitGroup

		for _, p := range profiles {
			wg.Add(1)
			go func(wg *sync.WaitGroup, p *pb.MatchProfile) {
				defer wg.Done()
				doFetchMatches(be, p, matchesForAssignment)
			}(&wg, p)
		}

		// Wait for all profiles to complete before proceeding.
		wg.Wait()
		telemetry.RecordUnitMeasurement(context.Background(), mIterations)
	}
}

func doFetchMatches(be pb.BackendServiceClient, profile *pb.MatchProfile, matchesForAssignment chan<- *pb.Match) {
	ctx, span := trace.StartSpan(context.Background(), "nbanow.director/FetchMatches")
	defer span.End()

	req := &pb.FetchMatchesRequest{
		Config: &pb.FunctionConfig{
			Host: "tasks.om-function",
			Port: 47102,
			Type: pb.FunctionConfig_GRPC,
		},
		Profile: profile,
	}

	telemetry.RecordUnitMeasurement(ctx, mFetchMatchCalls)

	stream, err := be.FetchMatches(ctx, req)
	if err != nil {
		telemetry.RecordUnitMeasurement(ctx, mFetchMatchErrors)
		logger.WithError(err).Error("failed to get available stream client")
		return
	}

	for {
		//logger.Infof("[Profile:%v] Wait..", profile.GetName())
		// BackEnd서버에서 매치정보를 읽는다.
		resp, err := stream.Recv()
		if err == io.EOF {
			telemetry.RecordUnitMeasurement(ctx, mFetchMatchSuccesses)
			return
		}

		if err != nil {
			logger.WithError(err).Error("failed to get matches from stream client")
			return
		}

		telemetry.RecordNUnitMeasurement(ctx, mSumTicketsReturned, int64(len(resp.GetMatch().Tickets)))
		telemetry.RecordUnitMeasurement(ctx, mMatchesReturned)

		//logger.Infof("[Profile:%v] GetMatchId(%v)", profile.GetName(), resp.GetMatch().MatchId)
		matchesForAssignment <- resp.GetMatch()
	}
}

func doAssignments(be pb.BackendServiceClient, matchesForAssignment <-chan *pb.Match) {
	ctx := context.Background()
	for m := range matchesForAssignment {
		ticketIds := make([]string, 0)
		for _, t := range m.GetTickets() {
			ticketIds = append(ticketIds, t.Id)
		}

		// TODO: 실시간서버의 IP로 대체
		address := fmt.Sprintf("%d.%d.%d.%d:2222", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))

		// Ticket에 할당할 Assign 정보
		req := &pb.AssignTicketsRequest{
			Assignments: []*pb.AssignmentGroup{
				{
					TicketIds: ticketIds,
					Assignment: &pb.Assignment{
						Connection: address,
					},
				},
			},
		}

		// Ticket에 할당할 Assign정보를 BackEnd서버로 보낸다
		if _, err := be.AssignTickets(ctx, req); err != nil {
			telemetry.RecordUnitMeasurement(ctx, mMatchAssignsFailed)
			logger.WithError(err).Errorf("AssignTickets failed for matchId(%v)", m.GetMatchId())
			continue
		}

		telemetry.RecordUnitMeasurement(ctx, mMatchesAssigned)
		logger.Infof("Assigned server %v to match %v", address, m.GetMatchId())
	}
}
