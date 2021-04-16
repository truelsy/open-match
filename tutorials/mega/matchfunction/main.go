package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"ns-open-match/pkg/matchfunction"
	"ns-open-match/pkg/pb"
	"sort"
	"time"
)

const matchName              = "basic-matchfunction"
const ticketsPerPoolPerMatch = 2

type MatchFunctionService struct {
	grpc               *grpc.Server
	queryServiceClient pb.QueryServiceClient
	port               int
}

func main() {
	// Connect to QueryService.
	conn, err := grpc.Dial("localhost:47103", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Open Match, got %s", err.Error())
	}
	defer conn.Close()

	mmfService := MatchFunctionService{
		queryServiceClient: pb.NewQueryServiceClient(conn),
	}

	// Create and host a new gRPC service on the configured port.
	server := grpc.NewServer()
	pb.RegisterMatchFunctionServer(server, &mmfService)

	serverPort := 50502
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatalf("TCP net listener initialization failed for port %v, got %s", serverPort, err.Error())
	}

	log.Printf("TCP net listener initialized for port %v", serverPort)
	err = server.Serve(ln)
	if err != nil {
		log.Fatalf("gRPC serve failed, got %s", err.Error())
	}
}

// Run is this match function's implementation of the gRPC call defined in api/matchfunction.proto.
func (s *MatchFunctionService) Run(req *pb.RunRequest, stream pb.MatchFunction_RunServer) error {
	// Fetch tickets for the pools specified in the Match Profile.
	log.Printf("Generating proposals for function %v", req.GetProfile().GetName())

	poolTickets, err := matchfunction.QueryPools(stream.Context(), s.queryServiceClient, req.GetProfile().GetPools())
	if err != nil {
		log.Printf("Failed to query tickets for the given pools, got %s", err.Error())
		return err
	}

	for pool, tickets := range poolTickets {
		log.Printf("Pool(%v) Tickets(%v)\n", pool, tickets)
	}

	// Generate proposals.
	proposals, err := makeMatches(req.GetProfile(), poolTickets)
	if err != nil {
		log.Printf("Failed to generate matches, got %s", err.Error())
		return err
	}

	//log.Printf("Streaming %v proposals to Open Match", len(proposals))
	for _, match := range proposals {
		log.Printf("Id(%v) User1(%v) User2(%v)\n", match.MatchId, match.Tickets[0].SearchFields.DoubleArgs["ovr"], match.Tickets[1].SearchFields.DoubleArgs["ovr"])
	}

	// Stream the generated proposals back to Open Match.
	for _, proposal := range proposals {
		log.Printf("Proposal : %+v\n", *proposal)
		if err := stream.Send(&pb.RunResponse{Proposal: proposal}); err != nil {
			log.Printf("Failed to stream proposals to Open Match, got %s", err.Error())
			return err
		}
	}

	return nil
}

func makeMatches(p *pb.MatchProfile, poolTickets map[string][]*pb.Ticket) ([]*pb.Match, error) {
	var matches []*pb.Match
	count := 0
	//for {
	//	insufficientTickets := false
	//	matchTickets := []*pb.Ticket{}
	//	for pool, tickets := range poolTickets {
	//		if len(tickets) < ticketsPerPoolPerMatch {
	//			// This pool is completely drained out. Stop creating matches.
	//			insufficientTickets = true
	//			break
	//		}
	//
	//		// Remove the Tickets from this pool and add to the match proposal.
	//		matchTickets = append(matchTickets, tickets[0:ticketsPerPoolPerMatch]...)
	//		poolTickets[pool] = tickets[ticketsPerPoolPerMatch:]
	//	}
	//
	//	if insufficientTickets {
	//		break
	//	}
	//
	//	matches = append(matches, &pb.Match{
	//		MatchId:       fmt.Sprintf("profile-%v-time-%v-%v", p.GetName(), time.Now().Format("2006-01-02T15:04:05.00"), count),
	//		MatchProfile:  p.GetName(),
	//		MatchFunction: matchName,
	//		Tickets:       matchTickets,
	//	})
	//
	//	count++
	//}

	//for {
	//	insufficientTickets := false
	//	matchTickets := make([]*pb.Ticket, 0)
	//	for pool, tickets := range poolTickets {
	//		if len(tickets) < ticketsPerPoolPerMatch {
	//			// This pool is completely drained out. Stop creating matches.
	//			insufficientTickets = true
	//			break
	//		}
	//
	//		curTicket := tickets[0]
	//		curOvr := curTicket.SearchFields.DoubleArgs["ovr"]
	//		rangeS := curOvr - (curOvr * 0.1)
	//		rangeE := curOvr + (curOvr * 0.1)
	//
	//		for i := 1; i < len(tickets); i++ {
	//			oppTicket := tickets[i]
	//			oppOvr := oppTicket.SearchFields.DoubleArgs["ovr"]
	//			if rangeS <= oppOvr && oppOvr <= rangeE {
	//				matchTickets = append(matchTickets, curTicket)
	//				matchTickets = append(matchTickets, oppTicket)
	//				poolTickets[pool] = append(tickets[:0], tickets[1:]...)
	//				poolTickets[pool] = append(tickets[:i-1], tickets[i:]...)
	//				break
	//			}
	//		}
	//	}
	//
	//	if insufficientTickets {
	//		break
	//	}
	//
	//	matches = append(matches, &pb.Match{
	//		MatchId:       fmt.Sprintf("profile-%v-time-%v-%v", p.GetName(), time.Now().Format("2006-01-02T15:04:05.00"), count),
	//		MatchProfile:  p.GetName(),
	//		MatchFunction: matchName,
	//		Tickets:       matchTickets,
	//	})
	//
	//	count++
	//}

	for {
		insufficientTickets := false

		for pool, tickets := range poolTickets {
			if len(tickets) < ticketsPerPoolPerMatch {
				// This pool is completely drained out. Stop creating matches.
				insufficientTickets = true
				break
			}

			sort.Slice(tickets, func(i, j int) bool {
				return tickets[i].SearchFields.DoubleArgs["ovr"] < tickets[j].SearchFields.DoubleArgs["ovr"]
			})

			for {
				if len(tickets) < ticketsPerPoolPerMatch {
					break
				}

				curTicket := tickets[0]
				nxtTicket := tickets[1]

				if curTicket.SearchFields.DoubleArgs["srange"] <= nxtTicket.SearchFields.DoubleArgs["ovr"] &&
					nxtTicket.SearchFields.DoubleArgs["ovr"] <= curTicket.SearchFields.DoubleArgs["erange"] {
					matches = append(matches, &pb.Match{
						MatchId:       fmt.Sprintf("profile-%v-time-%v-%v", p.GetName(), time.Now().Format("2006-01-02T15:04:05.00"), count),
						MatchProfile:  p.GetName(),
						MatchFunction: matchName,
						Tickets:       []*pb.Ticket{curTicket, nxtTicket},
					})
					count++
					tickets = tickets[ticketsPerPoolPerMatch:]
				} else {
					tickets = append(tickets[:0], tickets[1:]...)
				}

				poolTickets[pool] = tickets
			}
		}

		if insufficientTickets {
			break
		}
	}

	return matches, nil
}
