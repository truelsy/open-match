package main

import (
	"context"
	"math/rand"
	"ns-open-match/pkg/pb"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"log"
)

const (
	ticketsPerIter = 20
)

func isContextDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 200; i++ {
		go func() {
			for !isContextDone(ctx) {
				run(ctx)
				time.Sleep(5 * time.Second)
			}
		}()
	}

	quit := make(chan os.Signal)
	signal.Ignore(syscall.SIGPIPE)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit

	go func() {
		cancel()
	}()

}

// Ticket generates a Ticket with data using the package configuration.
func makeTicket(myOvr int) *pb.Ticket {
	//modes := []string{"mode.demo", "mode.ctf", "mode.battleroyale"}
	modes := []string{"arena"}
	ticket := &pb.Ticket{
		SearchFields: &pb.SearchFields{
			DoubleArgs: map[string]float64{
				"ovr":    float64(myOvr),
				"srange": float64(myOvr) - (float64(myOvr) * 0.1),
				"erange": float64(myOvr) + (float64(myOvr) * 0.1),
			},
			StringArgs: nil,
			Tags: []string{
				modes[rand.Intn(len(modes))],
			},
		},
	}
	return ticket
}

func run(ctx context.Context) {
	conn, err := grpc.Dial("localhost:47104", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Open Match, got %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	fe := pb.NewFrontendServiceClient(conn)

	low := 600
	high := 1000
	ovr := rand.Int63n(int64(high-low+1)) + int64(low)

	req := &pb.CreateTicketRequest{
		Ticket: makeTicket(int(ovr)),
	}

	resp, err := fe.CreateTicket(ctx, req)
	if err != nil {
		log.Fatalf("Failed to Create Ticket, got %s", err.Error())
	}

	for {
		got, err := fe.GetTicket(ctx, &pb.GetTicketRequest{TicketId: resp.GetId()})
		if err != nil {
			log.Fatalf("Failed to Get Ticket %v, got %s", resp.GetId(), err.Error())
		}

		//log.Printf("wait.... got(%+v)\n", got)

		if got.GetAssignment() != nil {
			log.Printf("Id(%v) Ovr(%v) Assignment(%v)", got.GetId(), got.SearchFields.DoubleArgs["ovr"], got.GetAssignment())
			break
		}

		time.Sleep(time.Second * 1)
	}

	_, err = fe.DeleteTicket(ctx, &pb.DeleteTicketRequest{TicketId: resp.GetId()})
	if err != nil {
		log.Fatalf("Failed to Delete Ticket %v, got %s", resp.GetId(), err.Error())
	}
}
