package greeter

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type useCase struct {
	repo IRepo
	log  *log.Helper
}

func NewUseCase(repo IRepo, logger log.Logger) IUseCase {
	return &useCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *useCase) SayHello(ctx context.Context, name string) error {
	return uc.repo.Create(ctx, &Entity{Name: name})
}
