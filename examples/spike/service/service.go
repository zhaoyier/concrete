package service

import (
	"context"

	"git.ezbuy.me/ezbuy/concrete/rpc"
)

type Hello struct {
}

func (s *Hello) SayHello(ctx context.Context, req *rpc.HelloRequest) (*rpc.HelloReply, error) {
	return &rpc.HelloReply{
		Message: "response",
	}, nil
}