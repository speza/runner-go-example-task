package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/speza/runner/proto"
	"time"
)

type Task struct {
}

func (e Task) Do(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	log.Infof("called execute")
	log.Info(request.Args)
	log.Info("sleeping to simulate long running task...")
	time.Sleep(4 * time.Second)
	log.Info("ok done now!")
	return &proto.Response{
		Message: "hello from dockerised task...",
	}, nil
}
