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

// Package evaluator provides the Evaluator service for Open Match golang harness.
package evaluator

import (
	"context"
	"io"

	"github.com/pkg/errors"
	"go.opencensus.io/stats"
	"golang.org/x/sync/errgroup"
	"ns-open-match/pkg/pb"
)

// Evaluator is the function signature for the Evaluator to be implemented by
// the user. The harness will pass the Matches to evaluate to the Evaluator
// and the Evaluator will return an accepted list of Matches.
type Evaluator func(ctx context.Context, in <-chan *pb.Match, out chan<- string) error

// evaluatorService implements pb.EvaluatorServer, the server generated by
// compiling the protobuf, by fulfilling the pb.EvaluatorServer interface.
type evaluatorService struct {
	evaluate Evaluator
}

// Evaluate is this harness's implementation of the gRPC call defined in
// api/evaluator.proto.
func (s *evaluatorService) Evaluate(stream pb.Evaluator_EvaluateServer) error {
	g, ctx := errgroup.WithContext(stream.Context())

	in := make(chan *pb.Match)
	out := make(chan string)

	g.Go(func() error {
		defer close(in)
		count := 0
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			select {
			case in <- req.Match:
				count++
			case <-ctx.Done():
				return ctx.Err()
			}
		}
		stats.Record(ctx, matchesPerEvaluateRequest.M(int64(count)))
		return nil
	})
	g.Go(func() error {
		defer close(out)
		return s.evaluate(ctx, in, out)
	})
	g.Go(func() error {
		defer func() {
			for range out {
			}
		}()

		count := 0
		for id := range out {
			err := stream.Send(&pb.EvaluateResponse{MatchId: id})
			if err != nil {
				return err
			}
			count++
		}
		stats.Record(ctx, matchesPerEvaluateResponse.M(int64(count)))
		return nil
	})

	err := g.Wait()
	return errors.Wrap(err, "Error in evaluator.Evaluate")
}
