package checker

import (
	"github.com/labstack/echo/v4"
	"interaction-api/domain/checker"
	"net/http"
)

type Controller struct {
	service checker.Service
}

func NewController(service checker.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) TestChecker(ctx echo.Context) error {
	m := map[string]bool{}

	if err := c.service.CheckDBService(); err != nil {
		m["DB Connection"] = false
	}
	m["DB Connection"] = true

	if err := c.service.CheckMongoDB(ctx.Request().Context()); err != nil {
		m["DB Mongo"] = false
	}
	m["DB Mongo"] = true

	return ctx.JSON(http.StatusOK, m)
}
