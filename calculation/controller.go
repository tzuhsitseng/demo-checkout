package calculation

import (
	"errors"

	"github.com/kataras/iris/v12"

	"github.com/tzuhsitseng/demo-checkout/domain"
)

type Controller struct {
	svc domain.CalcService
}

func NewController(svc domain.CalcService) *Controller {
	return &Controller{svc: svc}
}

func (c *Controller) Calculate(ctx iris.Context) {
	userLevel, err := ctx.URLParamInt("user_level")
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	if userLevel < 0 || userLevel > 3 {
		ctx.StopWithError(iris.StatusBadRequest, errors.New("invalid user level"))
		return
	}

	userPoints, err := ctx.URLParamInt("user_points")
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	if userPoints < 0 {
		ctx.StopWithError(iris.StatusBadRequest, errors.New("invalid user points"))
		return
	}

	price, err := ctx.URLParamInt("price")
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	if price < 0 {
		ctx.StopWithError(iris.StatusBadRequest, errors.New("invalid price"))
		return
	}

	coin, points, err := c.svc.Calculate(userLevel, userPoints, price)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	resp := domain.CalcResponse{
		Coin:   coin,
		Points: points,
	}
	_, _ = ctx.JSON(resp)
}
