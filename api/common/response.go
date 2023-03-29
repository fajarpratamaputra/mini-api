package common

import (
	"github.com/labstack/echo/v4"
	"interaction-api/config"
)

type Response struct {
	Data     interface{} `json:"data"`
	MetaData MetaData    `json:"meta"`
	Status   Status      `json:"status"`
}

type ResponseCreated struct {
	Status Status `json:"status"`
}

func CreateResponse(c echo.Context, httpStatus int, statusCode MessageEnum, data interface{}, pagination *Pagination) error {
	response := new(Response)
	metadata := MetaData{VideoPath: config.AppConfig.VideoPath, ImagePath: config.AppConfig.ImagePath}
	metadata.Pagination = pagination

	status := Status{}

	response.Data = data
	status.Code = statusCode.Int()
	status.MessageServer = statusCode.String()
	status.MessageClient = statusCode.String()

	response.Status = status
	response.MetaData = metadata
	return c.JSON(httpStatus, response)
}

func CreateResponseCreated(ctx echo.Context, httpStatus int, statusCode MessageEnum) error {
	response := new(ResponseCreated)

	status := Status{}

	status.Code = statusCode.Int()
	status.MessageServer = statusCode.String()
	status.MessageClient = statusCode.String()

	response.Status = status
	return ctx.JSON(httpStatus, response)
}
