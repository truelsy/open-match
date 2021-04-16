package main

import (
	"context"
	"flag"
	"math/rand"
	"ns-open-match/pkg/pb"
	"time"

	"google.golang.org/grpc"

	"log"
)

const (
	ticketsPerIter = 20
)

func main() {
	myOvr := flag.Int("ovr", 100, "OVR")
	flag.Parse()

	conn, err := grpc.Dial("localhost:47104", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Open Match, got %v", err)
	}
	defer conn.Close()

	//fe := pb.NewFrontendServiceClient(conn)
	//for range time.Tick(time.Second * 2) {
	//	for i := 0; i <= ticketsPerIter; i++ {
	//		req := &pb.CreateTicketRequest{
	//			Ticket: makeTicket(),
	//		}
	//
	//		resp, err := fe.CreateTicket(context.Background(), req)
	//		if err != nil {
	//			log.Printf("Failed to Create Ticket, got %s", err.Error())
	//			continue
	//		}
	//
	//		log.Println("Ticket created successfully, id:", resp.Id)
	//		go deleteOnAssign(fe, resp)
	//	}
	//}

	fe := pb.NewFrontendServiceClient(conn)

	req := &pb.CreateTicketRequest{
		Ticket: makeTicket(*myOvr),
	}

	resp, err := fe.CreateTicket(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to Create Ticket, got %s", err.Error())
	}

	for {
		got, err := fe.GetTicket(context.Background(), &pb.GetTicketRequest{TicketId: resp.GetId()})
		if err != nil {
			log.Fatalf("Failed to Get Ticket %v, got %s", resp.GetId(), err.Error())
		}

		log.Printf("wait.... got(%+v)\n", got)

		if got.GetAssignment() != nil {
			log.Printf("Ticket %v got assignment %v", got.GetId(), got.GetAssignment())
			break
		}

		time.Sleep(time.Second * 1)
	}

	_, err = fe.DeleteTicket(context.Background(), &pb.DeleteTicketRequest{TicketId: resp.GetId()})
	if err != nil {
		log.Fatalf("Failed to Delete Ticket %v, got %s", resp.GetId(), err.Error())
	}
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

// deleteOnAssign fetches the Ticket state periodically and deletes the Ticket
// once it has an assignment.
func deleteOnAssign(fe pb.FrontendServiceClient, t *pb.Ticket) {
	for {
		got, err := fe.GetTicket(context.Background(), &pb.GetTicketRequest{TicketId: t.GetId()})
		if err != nil {
			log.Fatalf("Failed to Get Ticket %v, got %s", t.GetId(), err.Error())
		}

		if got.GetAssignment() != nil {
			log.Printf("Ticket %v got assignment %v", got.GetId(), got.GetAssignment())
			break
		}

		time.Sleep(time.Second * 1)
	}

	_, err := fe.DeleteTicket(context.Background(), &pb.DeleteTicketRequest{TicketId: t.GetId()})
	if err != nil {
		log.Fatalf("Failed to Delete Ticket %v, got %s", t.GetId(), err.Error())
	}
}
