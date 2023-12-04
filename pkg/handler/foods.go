package handler

import (
	"fmt"
	"net/http"
	"sportnotes"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getAllFoodsResponse struct {
	Data []sportnotes.Food `json:"data"`
}

func (h *Handler) createFood(c *gin.Context) {
	// _, err := getDoctorId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	var input sportnotes.Food
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(input)

	id, err := h.services.FoodList.CreateFood(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newFoodResponse{
		Id: id,
	})
}

func (h *Handler) getAllFoods(c *gin.Context) {
	foods, err := h.services.FoodList.GetAllFoods()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllFoodsResponse{
		Data: foods,
	})
}

func (h *Handler) getFoodById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	food, err := h.services.FoodList.GetFoodById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, food)
}

func (h *Handler) updateFood(c *gin.Context) {
	// _, err := getDoctorId(c)
	// if err != nil {
	// 	return
	// }

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input sportnotes.UpdTraining
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TrainingList.UpdateTraining(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteFood(c *gin.Context) {
	// _, err := getDoctorId(c)
	// if err != nil {
	// 	return
	// }

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TrainingList.DeleteTraining(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
