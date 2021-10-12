package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/kratos-multi-repo/app-layout/internal/biz"
	v1 "github.com/kratos-multi-repo/helloworld-sdk/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

func (s *GreeterService) Ping(context.Context, *emptypb.Empty) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: "pong"}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
