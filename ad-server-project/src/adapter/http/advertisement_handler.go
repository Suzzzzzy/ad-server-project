package http

import (
	"ad-server-project/src/domain"
	"ad-server-project/src/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

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
	}
}

func (a *AdvertisementHandler) GetByCountryAndGender(c *gin.Context) {
	id := c.Query("id")
	userId, _ := strconv.Atoi(id)
	userGender := c.Query("gender")
	userCountry := c.Query("country")
	ctx := c.Request.Context()

	user := &model.User{
		ID:      userId,
		Gender:  userGender,
		Country: userCountry,
	}

	result, err := a.AdvertisementUsecase.GetByCountryAndGender(ctx, user)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
