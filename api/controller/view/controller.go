package view

import (
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
	"interaction-api/api/common"
	"interaction-api/domain"
	"interaction-api/domain/view"
	"net/http"
	"strconv"
)

type Controller struct {
	service view.Service
}

func NewController(service view.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) AddView(ctx echo.Context) error {
	userData := common.GetUserData(ctx)
	userId := int(userData["vid"].(float64))
	deviceId := userData["device_id"].(string)

	var iView domain.InteractionView
	rules := govalidator.MapData{
		"service":      []string{"required", "alpha_space"},
		"content_type": []string{"required", "alpha_space"},
		"content_id":   []string{"required", "numeric"},
	}
	validate := common.ValidateRequestPayload(ctx, rules, &iView)
	if validate != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.InvalidInput)
	}
	iView.UserID = userId
	iView.Action = "view"
	iView.DeviceID = deviceId

	if err := c.service.View(ctx.Request().Context(), iView); err != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.ServerError)
	}
	return common.CreateResponseCreated(ctx, http.StatusOK, common.Success)
}

func (c *Controller) GetTotalView(ctx echo.Context) error {
	v := domain.RequestGet{
		Service:     ctx.Param("service"),
		ContentType: ctx.Param("type"),
	}
	cId, _ := strconv.Atoi(ctx.Param("id"))

	v.ContentId = cId

	totalView, err := c.service.TotalView(ctx.Request().Context(), v)
	if err != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.NotFound)
	}

	data := domain.Total{
		Total:       totalView,
		TotalString: common.TotalFormatting(totalView),
	}

	return common.CreateResponse(ctx, http.StatusOK, common.Success, data, nil)
}
