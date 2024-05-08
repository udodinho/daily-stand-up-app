package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/udodinho/daily-standup-app/http/controller"
	"github.com/udodinho/daily-standup-app/http/server"
	gomock "go.uber.org/mock/gomock"
)

func TestCreateUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	controller := controller.Controller{Db: m}
	s := server.Server{Controller: &controller}
	router := s.SetupRouter()

	type Sprint string

	const (
		SprintOne Sprint = "one"
	)

	type Status string

	const (
		WithinStandup Status = "within standup"
	)

	payload := &struct {
		EmployeeID        int    `json:"employee_id" binding:"required"`
		EmployeeName      string `json:"employee_name" binding:"required"`
		Sprint            Sprint `json:"sprint" binding:"required"`
		TaskID            string `json:"task_id" binding:"required"`
		WorkDoneYesterday string `json:"work_done_yesterday"  binding:"required"`
		WorkToDoToday     string `json:"work_to_do_today" binding:"required"`
		BlockedBy         string `json:"blocked_by" binding:"required"`
		Breakaway         string `json:"breakaway" binding:"required"`
		Status            Status `json:"status" binding:"required"`
		Date              string `json:"date" binding:"required"`
	}{
		EmployeeID:        1,
		EmployeeName:      "John",
		Sprint:            SprintOne,
		TaskID:            "avg-234",
		WorkDoneYesterday: "Did work",
		WorkToDoToday:     "Will do work",
		BlockedBy:         "James",
		Breakaway:         "Made research",
		Status:            WithinStandup,
		Date:              "2024-02-08",
	}

	m.EXPECT().CreateUpdate(gomock.Any()).Return(nil, nil)

	jsonFile, err := json.Marshal(payload)
	if err != nil {
		t.Error("Failed to marshal file")
	}
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/checkin", strings.NewReader(string(jsonFile)))

	router.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusCreated, rw.Code)

}
