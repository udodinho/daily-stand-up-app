package domain

import "github.com/udodinho/daily-standup-app/entity"

// DB is an interface
type DB interface {
	CreateUpdate(checkin *entity.CheckIn) (*entity.CheckIn, error)
	FindAll(startWeek, endWeek, sprint, date, owner string) ([]entity.CheckIn, error)
}
