package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/udodinho/daily-standup-app/http/response"
)

func (c *Controller) GetAllCheckIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		weekStart := ctx.Query("startWeek")
		weekEnd := ctx.Query("endWeek")
		sprint := ctx.Query("sprint")
		day := ctx.Query("date")
		owner := ctx.Query("owner")

		updates, err := c.Db.FindAll(weekStart, weekEnd, sprint, day, owner)

		if err != nil {
			response.Failure(400, ctx, "Couldn't get all updates", true)
			return
		}

		response.Success(200, updates, ctx, "Successfully fetched all updates", false)
	}
}
