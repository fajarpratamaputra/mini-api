package follow

import (
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
	"interaction-api/api/common"
	"interaction-api/domain"
	"interaction-api/domain/follow"
	"net/http"
	"strconv"
)

type Controller struct {
	service follow.Service
}

func NewController(service follow.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) AddFollow(ctx echo.Context) error {
	userData := common.GetUserData(ctx)
	userId := int(userData["vid"].(float64))

	var iView domain.InteractionFollow
	rules := govalidator.MapData{
		"follow_to": []string{"required", "alpha_space"},
	}
	validate := common.ValidateRequestPayload(ctx, rules, &iView)
	if validate != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.InvalidInput)
	}
	iView.UserID = userId

	if err := c.service.Follow(ctx.Request().Context(), iView); err != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.ServerError)
	}
	return common.CreateResponseCreated(ctx, http.StatusOK, common.Success)
}

func (c *Controller) UnFollow(ctx echo.Context) error {
	userData := common.GetUserData(ctx)
	userId := int(userData["vid"].(float64))

	var iView domain.InteractionFollow
	rules := govalidator.MapData{
		"follow_to": []string{"required", "alpha_space"},
	}
	validate := common.ValidateRequestPayload(ctx, rules, &iView)
	if validate != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.InvalidInput)
	}
	iView.UserID = userId

	if err := c.service.UnFollow(ctx.Request().Context(), iView); err != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.ServerError)
	}
	return common.CreateResponseCreated(ctx, http.StatusOK, common.Success)
}

func (c *Controller) StatusFollow(ctx echo.Context) error {
	userData := common.GetUserData(ctx)
	userId := int(userData["vid"].(float64))
	followId, _ := strconv.Atoi(ctx.Param("id"))

	isFollow, err := c.service.GetStatusFollow(ctx.Request().Context(), userId, followId)
	if err != nil {
		return err
	}

	data := map[string]bool{}
	data["is_follow"] = isFollow

	return common.CreateResponse(ctx, http.StatusOK, common.Success, data, nil)
}

func (c *Controller) GetTotalFollow(ctx echo.Context) error {
	cId, _ := strconv.Atoi(ctx.Param("id"))

	following := domain.RequestGetFollow{
		UserID: cId,
	}

	totalFollowing, err := c.service.TotalFollow(ctx.Request().Context(), following)
	if err != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.NotFound)
	}

	follower := domain.RequestGetFollower{
		FollowTo: cId,
	}

	totalFollower, err := c.service.TotalFollower(ctx.Request().Context(), follower)
	if err != nil {
		return common.CreateResponseCreated(ctx, http.StatusOK, common.NotFound)
	}

	data := domain.TotalFollow{
		TotalFollowing:       totalFollowing,
		TotalStringFollowing: common.TotalFormatting(totalFollowing),
		TotalFollower:        totalFollower,
		TotalStringFollower:  common.TotalFormatting(totalFollower),
	}

	return common.CreateResponse(ctx, http.StatusOK, common.Success, data, nil)
}
