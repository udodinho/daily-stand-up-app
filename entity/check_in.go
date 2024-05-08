package entity

import (
	"time"

	"github.com/google/uuid"
)

type Sprint string

const (
	SprintOne   Sprint = "one"
	SprintTwo   Sprint = "two"
	SprintThree Sprint = "three"
	SprintFour  Sprint = "four"
	SprintFive  Sprint = "five"
)

type Status string

const (
	BeforeStandup Status = "before standup"
	AfterStandup  Status = "after standup"
	WithinStandup Status = "within standup"
)

type CheckIn struct {
	ID                uuid.UUID `gorm:"primarykey" json:"id"`
	EmployeeID        int       `json:"employee_id" binding:"required"`
	EmployeeName      string    `json:"employee_name" binding:"required"`
	Sprint            Sprint    `json:"sprint" binding:"required"`
	TaskID            string    `json:"task_id" binding:"required"`
	WorkDoneYesterday string    `json:"work_done_yesterday"  binding:"required"`
	WorkToDoToday     string    `json:"work_to_do_today" binding:"required"`
	BlockedBy         string    `json:"blocked_by" binding:"required"`
	Breakaway         string    `json:"breakaway" binding:"required"`
	CheckIn           time.Time `json:"check_in"`
	Status            Status    `json:"status" binding:"required"`
	Date              string    `json:"date" binding:"required"`
}
