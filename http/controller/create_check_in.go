package controller

import (
	// "encoding/binary"
	"fmt"
	// "math/big"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/udodinho/daily-standup-app/entity"
	"github.com/udodinho/daily-standup-app/http/response"
)

type Employees struct {
	EmployeeID   int
	EmployeeName string
}

func (c *Controller) CreateCheckIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employees := []Employees{
			{EmployeeID: 1, EmployeeName: "Honda"},
			{EmployeeID: 2, EmployeeName: "Kate"},
			{EmployeeID: 3, EmployeeName: "Faith"},
			{EmployeeID: 4, EmployeeName: "James"},
			{EmployeeID: 5, EmployeeName: "Kings"},
		}

		// Get a random index for selecting an employee
		randomIndex := rand.Intn(len(employees))

		// Get the random employee from the slice
		randomEmployee := employees[randomIndex]

		update := &entity.CheckIn{}

		date := time.Now()
		formattedDate := date.Format("2006-01-02")

		update.ID = uuid.New()
		update.EmployeeID = randomEmployee.EmployeeID
		update.EmployeeName = randomEmployee.EmployeeName
		update.CheckIn = time.Now()
		update.Date = formattedDate

		err := ctx.BindJSON(update)
		if err != nil {
			response.Failure(500, ctx, "unable to decode data, please provide all field", true)
			return
		}
		fmt.Println("New", update)

		_, err = c.Db.CreateUpdate(update)
		if err != nil {
			response.Failure(500, ctx, "unable to create update", true)
			return
		}

		response.Success(201, update, ctx, "successfully created an update", false)

	}
}
