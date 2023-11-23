package http

import (
	"ad-server-project/src/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


type AdvertisementHandler struct {
	AdvertisementUsecase model.AdvertisementUsecase
}

func NewAdvertisementHandler(r *gin.Engine, ad model.AdvertisementUsecase) {
	handler := &AdvertisementHandler{
		AdvertisementUsecase: ad,
	}
	router := r.Group("/ad-campaigns")
	{
		router.GET("", handler.GetByCountryAndGender)
		router.PUT("/reward", handler.UpdateReward)
	}
}

func (a *AdvertisementHandler) GetByCountryAndGender(c *gin.Context) {
	id := c.Query("user_id")
	userId, _ := strconv.Atoi(id)
	userGender := c.Query("user_gender")
	userCountry := c.Query("user_country")
	ctx := c.Request.Context()

	result, err := a.AdvertisementUsecase.GetByCountryAndGender(ctx, userId, userGender, userCountry)
	if err != nil {
		c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (a *AdvertisementHandler) UpdateReward(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: "Invalid ID"})
		return
	}
	reward, err := strconv.Atoi(c.Query("reward"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: "Invalid reward"})
		return
	}
	ctx := c.Request.Context()

	err = a.AdvertisementUsecase.UpdateReward(ctx, id, reward)
	if err != nil {
		c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}