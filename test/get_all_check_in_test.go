package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/udodinho/daily-standup-app/entity"
	"github.com/udodinho/daily-standup-app/http/controller"
	"github.com/udodinho/daily-standup-app/http/server"
	gomock "go.uber.org/mock/gomock"
)

func TestGetCheckins(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockDB(ctrl)
	controller := controller.Controller{Db: mockDB}
	s := server.Server{Controller: &controller}
	router := s.SetupRouter()

	uuidStr := "f9421eec-5c49-4c0d-b1e3-eafb2787579a"

	// Parse the string into a UUID
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
		return
	}

	data := &[]entity.CheckIn{
		{
			ID:                parsedUUID,
			EmployeeID:        1,
			EmployeeName:      "John",
			Sprint:            "one",
			TaskID:            "avg-234",
			WorkDoneYesterday: "Did work",
			WorkToDoToday:     "Will do work",
			BlockedBy:         "James",
			Breakaway:         "Made research",
			CheckIn:           time.Now(),
			Status:            "within standup",
			Date:              "2024-02-08",
		},
		{
			ID:                parsedUUID,
			EmployeeID:        1,
			EmployeeName:      "John",
			Sprint:            "one",
			TaskID:            "avg-234",
			WorkDoneYesterday: "Did work",
			WorkToDoToday:     "Will do work",
			BlockedBy:         "James",
			Breakaway:         "Made research",
			CheckIn:           time.Now(),
			Status:            "within standup",
			Date:              "2024-02-08",
		},
	}

	weekStart := ""
	weekEnd := ""
	sprint := "one"
	day := ""
	owner := ""

	mockDB.EXPECT().FindAll(weekStart, weekEnd, sprint, day, owner).Return(*data, nil)
	jsonFile, err := json.Marshal(data)
	if err != nil {
		t.Error("Failed to marshal file")
	}

	rw := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/checkins?sprint=one", nil)

	router.ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)
	assert.Contains(t, rw.Body.String(), string(jsonFile))
}
