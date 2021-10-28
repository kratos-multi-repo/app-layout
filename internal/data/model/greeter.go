package model

import "gorm.io/gorm"

type Greeter struct {
	gorm.Model

	Name string `gorm:"size:16"`
}

func (g *Greeter) TableName() string {
	return "t_greeters"
}
