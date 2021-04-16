package matchfunction

import (
	"fmt"
	"log"
	"net"
	utilTesting "ns-open-match/internal/util/testing"
	scenarios "ns-open-match/nbanow"
	"ns-open-match/pkg/pb"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	queryServiceAddr = "tasks.om-query:47103" // Address of the QueryService endpoint.
	serverPort       = 47102                  // The port for hosting the Match Function.
)

var (
	logger = logrus.WithFields(logrus.Fields{
		"app":       "openmatch",
		"component": "nbanow.matchfunction",
	})
)

func Run() {
	conn, err := grpc.Dial(queryServiceAddr, utilTesting.NewGRPCDialOptions(logger)...)
	if err != nil {
		log.Fatalf("Failed to connect to Open Match, got %v", err)
	}
	defer conn.Close()

	server := grpc.NewServer()

	// MatchFunction 등록
	pb.RegisterMatchFunctionServer(server, scenarios.ActiveScenario.MatchFunction)

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

	logger.Infof("TCP net listener initialized for port : %v", serverPort)
	err = server.Serve(ln)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("gRPC serve() error")
	}
}
