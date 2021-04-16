// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package statestore

import (
	"context"

	"go.opencensus.io/trace"
	"ns-open-match/pkg/pb"
)

// instrumentedService is a wrapper for a statestore service that provides instrumentation (metrics and tracing) of the database.
type instrumentedService struct {
	s Service
}

func (is *instrumentedService) Close() error {
	return is.s.Close()
}

func (is *instrumentedService) HealthCheck(ctx context.Context) error {
	err := is.s.HealthCheck(ctx)
	return err
}

func (is *instrumentedService) CreateTicket(ctx context.Context, ticket *pb.Ticket) error {
	ctx, span := trace.StartSpan(ctx, "statestore/instrumented.CreateTicket")
	defer span.End()
	return is.s.CreateTicket(ctx, ticket)
}

func (is *instrumentedService) GetTicket(ctx context.Context, id string) (*pb.Ticket, error) {
	ctx, span := trace.StartSpan(ctx, "statestore/instrumented.GetTicket")
	defer span.End()
	return is.s.GetTicket(ctx, id)
}

func (is *instrumentedService) DeleteTicket(ctx context.Context, id string) error {
	ctx, span := trace.StartSpan(ctx, "statestore/instrumented.DeleteTicket")
	defer span.End()
	return is.s.DeleteTicket(ctx, id)
}

func (is *instrumentedService) IndexTicket(ctx context.Context, ticket *pb.Ticket) error {
	ctx, span := trace.StartSpan(ctx, "statestore/instrumented.IndexTicket")
	defer span.End()
	return is.s.IndexTicket(ctx, ticket)
}

func (is *instrumentedService) DeindexTicket(ctx context.Context, id string) error {
	ctx, span := trace.StartSpan(ctx, "statestore/instrumented.DeindexTicket")
	defer span.End()
	return is.s.DeindexTicket(ctx, id)
}

func (is *instrumentedService) GetTickets(ctx context.Context, ids []string) ([]*pb.Ticket, error) {
	ctx, span := trace.StartSpan(ctx, "statestore/instrumented.GetTickets")
	defer span.End()
	return is.s.GetTickets(ctx, ids)
}

func (is *instrumentedService) GetIndexedIDSet(ctx context.Context) (map[string]struct{}, error) {
	ctx, span := trace.StartSpan(ctx, "statestore/instrumented.GetIndexedIDSet")
	defer span.End()
	return is.s.GetIndexedIDSet(ctx)
}

func (is *instrumentedService) UpdateAssignments(ctx context.Context, req *pb.AssignTicketsRequest) (*pb.AssignTicketsResponse, []*pb.Ticket, error) {
	ctx, span := trace.StartSpan(ctx, "statestore/instrumented.UpdateAssignments")
	defer span.End()
	return is.s.UpdateAssignments(ctx, req)
}

func (is *instrumentedService) GetAssignments(ctx context.Context, id string, callback func(*pb.Assignment) error) error {
	ctx, span := trace.StartSpan(ctx, "statestore/instrumented.GetAssignments")
	defer span.End()
	return is.s.GetAssignments(ctx, id, callback)
}

func (is *instrumentedService) AddTicketsToPendingRelease(ctx context.Context, ids []string) error {
	ctx, span := trace.StartSpan(ctx, "statestore/instrumented.AddTicketsToPendingRelease")
	defer span.End()
	return is.s.AddTicketsToPendingRelease(ctx, ids)
}

func (is *instrumentedService) DeleteTicketsFromPendingRelease(ctx context.Context, ids []string) error {
	ctx, span := trace.StartSpan(ctx, "statestore/instrumented.DeleteTicketsFromPendingRelease")
	defer span.End()
	return is.s.DeleteTicketsFromPendingRelease(ctx, ids)
}

func (is *instrumentedService) ReleaseAllTickets(ctx context.Context) error {
	ctx, span := trace.StartSpan(ctx, "statestore/instrumented.ReleaseAllTickets")
	defer span.End()
	return is.s.ReleaseAllTickets(ctx)
}
