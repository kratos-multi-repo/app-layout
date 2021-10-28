package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kratos-multi-repo/app-layout/internal/conf"
	"github.com/kratos-multi-repo/app-layout/internal/data/model"
	"github.com/kratos-multi-repo/pkg/dbutil"
	"gorm.io/gorm"
)

// Client .
type Client struct {
	// TODO wrapped database client
	platform *gorm.DB
}

// NewClient .
func NewClient(c *conf.Data, logger log.Logger) (*Client, func(), error) {
	log.NewHelper(logger).Infof("diver: %s, dsn: %s", c.Database.GetDriver(), c.Database.GetSource())

	client, err := dbutil.Connect(&dbutil.Option{
		Name:   "platform",
		Logger: logger,
		Config: &dbutil.ConnectConfig{
			Dsn:         c.Database.GetSource(),
			Driver:      c.Database.GetDriver(),
			MaxIdle:     300,
			MaxOpen:     300,
			LogMode:     true,
			MaxLifeTime: 300,
		},
	})
	if nil != err {
		return nil, nil, err
	}

	// todo: migrate
	if err := migrate(client); nil != err {
		return nil, nil, err
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the dbClient resources")
	}

	return &Client{
		platform: client,
	}, cleanup, nil
}

func migrate(platform *gorm.DB) error {
	return platform.AutoMigrate(
		&model.Greeter{},
	)
}
