package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/kratos-multi-repo/app-layout/internal/biz/greeter"
	"github.com/kratos-multi-repo/app-layout/internal/data/model"
	"github.com/save95/xerror"
	"github.com/save95/xerror/xcode"
	"gorm.io/gorm"
)

type greeterRepo struct {
	db  *gorm.DB
	log *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(client *Client, logger log.Logger) greeter.IRepo {
	return &greeterRepo{
		db:  client.platform,
		log: log.NewHelper(logger),
	}
}

func (r *greeterRepo) Create(ctx context.Context, g *greeter.Entity) error {
	r.log.Infof("%s say hello", g.Name)
	if len(g.Name) == 0 {
		return xerror.WithXCode(xcode.DBRequestParamError)
	}

	// storage
	entity := &model.Greeter{Name: g.Name}
	if err := r.db.WithContext(ctx).Create(entity).Error; nil != err {
		return xerror.WrapWithXCode(err, xcode.DBFailed)
	}

	return nil
}
