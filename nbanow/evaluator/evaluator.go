package evaluator

import (
	"fmt"
	"net"
	scenarios "ns-open-match/nbanow"
	"ns-open-match/pkg/pb"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	serverPort = 47108
)

var (
	logger = logrus.WithFields(logrus.Fields{
		"app":       "openmatch",
		"component": "nbanow.evaluator",
	})
)

func Run() {
	server := grpc.NewServer( /*utilTesting.NewGRPCServerOptions(logger)...*/ )
	pb.RegisterEvaluatorServer(server, scenarios.ActiveScenario.Evaluator)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"port":  serverPort,
		}).Fatal("net.Listen() error")
	}

	logger.WithFields(logrus.Fields{
		"port": serverPort,
	}).Info("TCP net listener initialized")

	logger.Info("TCP net listener initialized for port %v", serverPort)
	err = server.Serve(ln)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("gRPC serve() error")
	}
}
