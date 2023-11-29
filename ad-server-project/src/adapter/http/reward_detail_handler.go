package http

import (
	"ad-server-project/src/domain/model"
	"fmt"
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
		router.GET("/detail", handler.GetRecent)
		router.GET("/", handler.GetRewardBalance)
	}
}

type deductRewardRequest struct {
	Reward int `json:"reward"`
	UserId int `json:"user_id"`
}

type earnRewardRequest struct {
	AdId   int `json:"ad_id"`
	Reward int `json:"reward"`
	UserId int `json:"user_id"`
}

func (r *RewardDetailHandler) EarnReward(c *gin.Context) {
	var input earnRewardRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx := c.Request.Context()
	fmt.Printf("유저 아이디 %v", input.UserId)
	err := r.RewardDetailUsecase.EarnRewardDetail(ctx, input.AdId, input.Reward, input.UserId)
	if err != nil {
		c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Reward earned successfully")
}

func (r *RewardDetailHandler) DeductReward(c *gin.Context) {
	var input deductRewardRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx := c.Request.Context()

	err := r.RewardDetailUsecase.DeductRewardDetail(ctx, input.Reward, input.UserId)
	if err != nil {
		c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Reward deducted successfully")
}

func (r *RewardDetailHandler) GetRecent(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	ctx := c.Request.Context()

	result, err := r.RewardDetailUsecase.GetRecent(ctx, userId)
	if err != nil {
		c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (r *RewardDetailHandler) GetRewardBalance(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	ctx := c.Request.Context()

	result, err := r.RewardDetailUsecase.GetRewardBalance(ctx, userId)
	if err != nil {
		c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
