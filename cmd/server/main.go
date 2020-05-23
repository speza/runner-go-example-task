package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/speza/runner/proto"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Executor struct {
}

func (e Executor) Do(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	log.Infof("called execute")
	log.Info(request.Args)
	log.Info("sleeping to simulate long running task...")
	time.Sleep(4 * time.Second)
	log.Info("ok done now!")
	return &proto.Response{
		Message: "hello from dockerised task...",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":5300")
	if err != nil {
		log.WithError(err).Info("failed to listen")
		return
	}

	server := grpc.NewServer()

	proto.RegisterExecutorServer(server, &Executor{})

	log.Info("serving on :5300")
	if err := server.Serve(listen); err != nil {
		log.WithError(err).Info("failed to serve")
	}
}
