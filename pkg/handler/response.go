package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errMessage struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status" example:"ok"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errMessage{message})
}

type newUserResponse struct {
	Id int `json:"id" swaggertype:"primitive,integer"`
}

type newWorkoutResponse struct {
	Id int `json:"id" swaggertype:"primitive,integer"`
}

type newTreaningResponse struct {
	Id int `json:"id" swaggertype:"primitive,integer"`
}

type newNutritionResponse struct {
	Id int `json:"id" swaggertype:"primitive,integer"`
}
