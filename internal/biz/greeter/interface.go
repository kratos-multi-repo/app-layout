package greeter

import "context"

type IRepo interface {
	Create(context.Context, *Entity) error
}

type IUseCase interface {
	SayHello(ctx context.Context, name string) error
}
