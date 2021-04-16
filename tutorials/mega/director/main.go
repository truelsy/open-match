package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"ns-open-match/pkg/pb"
	"sync"
	"time"

	"context"
	"log"
	"math/rand"
)

func main() {
	conn, err := grpc.Dial("localhost:47105", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Open Match Backend, got %s", err.Error())
	}

	defer conn.Close()
	be := pb.NewBackendServiceClient(conn)

	// Generate the profiles to fetch matches for.
	profiles := generateProfiles()
	log.Printf("Fetching matches for %v profiles", len(profiles))

	for range time.Tick(time.Second * 5) {
		// Fetch matches for each profile and make random assignments for Tickets in
		// the matches returned.
		var wg sync.WaitGroup
		for _, p := range profiles {
			wg.Add(1)
			go func(wg *sync.WaitGroup, p *pb.MatchProfile) {
				defer wg.Done()
				matches, err := fetch(be, p)
				if err != nil {
					log.Printf("Failed to fetch matches for profile %v, got %s", p.GetName(), err.Error())
					return
				}

				log.Println(matches)

				log.Printf("Generated %v matches for profile %v", len(matches), p.GetName())
				if err := assign(be, matches); err != nil {
					log.Printf("Failed to assign servers to matches, got %s", err.Error())
					return
				}
			}(&wg, p)
		}

		wg.Wait()
	}
}

func generateProfiles() []*pb.MatchProfile {
	var profiles []*pb.MatchProfile
	//modes := []string{"mode.demo", "mode.ctf", "mode.battleroyale"}
	modes := []string{"mode.demo"}
	for _, mode := range modes {
		profiles = append(profiles, &pb.MatchProfile{
			Name: "mode_based_profile",
			Pools: []*pb.Pool{
				{
					Name: "pool_mode_" + mode,
					TagPresentFilters: []*pb.TagPresentFilter{
						{
							Tag: mode,
						},
					},
				},
			},
		})
	}
	return profiles
}

const (
	functionHostName = "172.17.0.1"
	functionPort     int32 = 50502
)

func fetch(be pb.BackendServiceClient, p *pb.MatchProfile) ([]*pb.Match, error) {
	req := &pb.FetchMatchesRequest{
		Config: &pb.FunctionConfig{
			Host: functionHostName,
			Port: functionPort,
			Type: pb.FunctionConfig_GRPC,
		},
		Profile: p,
	}

	stream, err := be.FetchMatches(context.Background(), req)
	if err != nil {
		log.Println()
		return nil, err
	}

	log.Printf("Profile(%v)\n", p.Name)

	var result []*pb.Match
	for {
		//log.Print("Wait....")
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		log.Printf("Resp : %+v\n", resp)

		if err != nil {
			return nil, err
		}

		result = append(result, resp.GetMatch())
	}
	return result, nil
}

func assign(be pb.BackendServiceClient, matches []*pb.Match) error {
	for _, match := range matches {
		ticketIDs := []string{}
		for _, t := range match.GetTickets() {
			ticketIDs = append(ticketIDs, t.Id)
		}

		conn := fmt.Sprintf("%d.%d.%d.%d:2222", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
		req := &pb.AssignTicketsRequest{
			Assignments: []*pb.AssignmentGroup{
				{
					TicketIds: ticketIDs,
					Assignment: &pb.Assignment{
						Connection: conn,
					},
				},
			},
		}

		if _, err := be.AssignTickets(context.Background(), req); err != nil {
			return fmt.Errorf("AssignTickets failed for match %v, got %w", match.GetMatchId(), err)
		}

		log.Printf("Assigned server %v to match %v", conn, match.GetMatchId())
	}

	return nil
}
