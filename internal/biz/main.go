package biz

import (
	"github.com/google/wire"
	"github.com/kratos-multi-repo/app-layout/internal/biz/greeter"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	greeter.NewUseCase,
)
