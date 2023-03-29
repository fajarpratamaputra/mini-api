package likes

import (
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
	"interaction-api/api/common"
	"interaction-api/domain"
	"interaction-api/domain/like"
	"net/http"
	"strconv"
)

type Controller struct {
	service like.Service
}

func NewController(service like.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) AddLike(ctx echo.Context) error {
	userData := common.GetUserData(ctx)
	userId := int(userData["vid"].(float64))

	var interaction domain.Interaction
	rules := govalidator.MapData{
		"service":      []string{"required", "alpha_space"},
		"content_type": []string{"required", "alpha_space"},
		"content_id":   []string{"required", "numeric"},
	}
	validate := common.ValidateRequestPayload(ctx, rules, &interaction)
	if validate != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.InvalidInput)
	}
	interaction.UserID = userId
	interaction.Action = "like"

	if err := c.service.Like(ctx.Request().Context(), interaction); err != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.ServerError)
	}
	return common.CreateResponseCreated(ctx, http.StatusOK, common.Success)
}

func (c *Controller) Unlike(ctx echo.Context) error {
	userData := common.GetUserData(ctx)
	userId := int(userData["vid"].(float64))

	var interaction domain.Interaction
	rules := govalidator.MapData{
		"service":      []string{"required", "alpha_space"},
		"content_type": []string{"required", "alpha_space"},
		"content_id":   []string{"required", "numeric"},
	}
	validate := common.ValidateRequestPayload(ctx, rules, &interaction)
	if validate != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.InvalidInput)
	}
	interaction.UserID = userId
	interaction.Action = "unlike"

	if err := c.service.Unlike(ctx.Request().Context(), interaction); err != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.ServerError)
	}
	return common.CreateResponseCreated(ctx, http.StatusOK, common.Success)
}

func (c *Controller) TotalLike(ctx echo.Context) error {
	v := domain.RequestGet{
		Service:     ctx.Param("service"),
		ContentType: ctx.Param("type"),
	}
	cId, _ := strconv.Atoi(ctx.Param("id"))

	v.ContentId = cId

	totalLike, err := c.service.TotalLike(ctx.Request().Context(), v)
	if err != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.NotFound)
	}

	data := domain.Total{
		Total:       totalLike,
		TotalString: common.TotalFormatting(totalLike),
	}

	return common.CreateResponse(ctx, http.StatusOK, common.Success, data, nil)
}
