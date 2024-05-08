package controller

import (
	"github.com/udodinho/daily-standup-app/database"
	"github.com/udodinho/daily-standup-app/domain"
)

type Controller struct {
	Db domain.DB
}

func NewController() *Controller {
	return &Controller{Db: database.NewDatabase()}
}
