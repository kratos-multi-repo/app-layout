package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/kratos-multi-repo/app-layout/internal/biz/greeter"
	v1 "github.com/kratos-multi-repo/helloworld-sdk/v1"
	"github.com/kratos-multi-repo/pkg/errutil"
	"github.com/save95/xerror"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  greeter.IUseCase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc greeter.IUseCase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

// Ping implements v1.GreeterServer and v1.GreeterHTTPServer
func (s *GreeterService) Ping(context.Context, *emptypb.Empty) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: "pong"}, nil
}

// SayHello implements v1.GreeterServer and v1.GreeterHTTPServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	// show error
	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	// show xerror
	if in.GetName() == "xerror" {
		return nil, errutil.Kratos(xerror.WithCode(1003, "xerror failed"))
	}

	if err := s.uc.SayHello(ctx, in.GetName()); nil != err {
		return nil, err
	}

	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
