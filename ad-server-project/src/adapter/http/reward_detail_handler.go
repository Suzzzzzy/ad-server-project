package http

import (
	"ad-server-project/src/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RewardDetailHandler struct {
	RewardDetailUsecase model.RewardDetailUsecase
}

func NewRewardDetailHandler(r *gin.Engine, re model.RewardDetailUsecase) {
	handler := &RewardDetailHandler{
		RewardDetailUsecase: re,
	}
	router := r.Group("/reward")
	{
		router.POST("/plus", handler.EarnReward)
		router.POST("/minus", handler.DeductReward)
	}
}


func (r *RewardDetailHandler) EarnReward(c *gin.Context) {
	adId, _ := strconv.Atoi(c.Query("ad_id"))
	reward, _ := strconv.Atoi(c.Query("reward"))
	userId , _ := strconv.Atoi(c.Query("user_id"))
	ctx := c.Request.Context()

	err := r.RewardDetailUsecase.EarnRewardDetail(ctx, adId, reward, userId)
	if err != nil {
		c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (r *RewardDetailHandler) DeductReward(c *gin.Context) {
	reward, _ := strconv.Atoi(c.Query("reward"))
	userId , _ := strconv.Atoi(c.Query("user_id"))
	ctx := c.Request.Context()

	err := r.RewardDetailUsecase.DeductRewardDetail(ctx, reward, userId)
	if err != nil {
		c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}