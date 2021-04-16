package battle

import (
	"fmt"
	"io"
	"ns-open-match/internal/config"
	"ns-open-match/pkg/pb"
	"sort"
	"time"
)

const (
	poolName = "battle"
)

type ScenarioBattle struct {
	modes []string
}

func Scenario() *ScenarioBattle {
	c, err := config.Read()
	if err != nil {
		panic("can't read configuration file")
	}

	return &ScenarioBattle{
		modes: c.GetStringSlice("scenario.battle.modes"),
	}
}

func (s *ScenarioBattle) Profiles() []*pb.MatchProfile {
	profiles := make([]*pb.MatchProfile, 0)
	for _, mode := range s.modes {
		profiles = append(profiles, &pb.MatchProfile{
			Name: fmt.Sprintf("profile_%v", mode),
			Pools: []*pb.Pool{
				{
					Name:              poolName + "_" + mode,
					TagPresentFilters: []*pb.TagPresentFilter{{Tag: mode}},
				},
			},
		})
	}
	return profiles
}

func (s *ScenarioBattle) MatchFunction(profile *pb.MatchProfile, poolTickets map[string][]*pb.Ticket) ([]*pb.Match, error) {
	switch profile.GetName() {
	case "profile_arena":
		return s.matchMakeArena(profile, poolTickets)
	}
	return nil, nil
}

func (s *ScenarioBattle) matchMakeArena(profile *pb.MatchProfile, poolTickets map[string][]*pb.Ticket) ([]*pb.Match, error) {
	matches := make([]*pb.Match, 0)
	matchCnt := 0
	matchUserCnt := 2

	for {
		insufficientTickets := false

		for pool, tickets := range poolTickets {
			if len(tickets) < matchUserCnt {
				// This pool is completely drained out. Stop creating matches.
				insufficientTickets = true
				break
			}

			sort.Slice(tickets, func(i, j int) bool {
				return tickets[i].SearchFields.DoubleArgs["ovr"] < tickets[j].SearchFields.DoubleArgs["ovr"]
			})

			for {
				if len(tickets) < matchUserCnt {
					break
				}

				curTicket := tickets[0]
				nxtTicket := tickets[1]

				if curTicket.SearchFields.DoubleArgs["srange"] <= nxtTicket.SearchFields.DoubleArgs["ovr"] &&
					nxtTicket.SearchFields.DoubleArgs["ovr"] <= curTicket.SearchFields.DoubleArgs["erange"] {
					matches = append(matches, &pb.Match{
						MatchId:       fmt.Sprintf("profile-%v-time-%v-%v", profile.GetName(), time.Now().Format("2006-01-02T15:04:05.00"), matchCnt),
						MatchProfile:  profile.GetName(),
						MatchFunction: "arena",
						Tickets:       []*pb.Ticket{curTicket, nxtTicket},
					})
					matchCnt++
					tickets = tickets[matchUserCnt:]
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

func (s *ScenarioBattle) Evaluate(stream pb.Evaluator_EvaluateServer) error {
	used := map[string]struct{}{}

	// TODO: once the evaluator client supports sending and recieving at the
	// same time, don't buffer, just send results immediately.
	matchIDs := []string{}

outer:
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading evaluator input stream: %w", err)
		}

		m := req.GetMatch()

		for _, t := range m.Tickets {
			if _, ok := used[t.Id]; ok {
				continue outer
			}
		}

		for _, t := range m.Tickets {
			used[t.Id] = struct{}{}
		}

		matchIDs = append(matchIDs, m.GetMatchId())
	}

	for _, mID := range matchIDs {
		err := stream.Send(&pb.EvaluateResponse{MatchId: mID})
		if err != nil {
			return fmt.Errorf("error sending evaluator output stream: %w", err)
		}
	}

	return nil
}
