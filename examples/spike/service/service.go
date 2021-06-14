package service

import (
	"context"
	"fmt"

	"git.ezbuy.me/ezbuy/concrete/rpc"
)

type Hello struct {
}

func (s *Hello) SayHello(ctx context.Context, req *rpc.HelloRequest) (*rpc.HelloReply, error) {
	fmt.Println("==>>TODO SayHello 002")

	return &rpc.HelloReply{
		Message: "response",
	}, nil
}
