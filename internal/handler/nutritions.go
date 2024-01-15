package handler

import (
	"net/http"
	"sportnotes/internal/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createNutrition(c *gin.Context) {
	idUser, err := strconv.Atoi(c.Param("idUser"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input schemas.Nutrition
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.NutritionList.CreateNutrition(idUser, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newIDResponse{
		ID: id,
	})
}

func (h *Handler) getNutritionsByParam(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	startPoint := c.DefaultQuery("startpoint", "none")
	if startPoint == "" {
		startPoint = "week"
	}

	endPoint := c.DefaultQuery("endpoint", "none")
	if endPoint == "" {
		endPoint = "week"
	}

	nutritions, err := h.services.NutritionList.GetNutritionsByParam(id, startPoint, endPoint)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nutritions)
}

// func (h *Handler) updateNutrition(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
// 		return
// 	}

// 	var input sportnotes.UpdWorkout
// 	if err := c.BindJSON(&input); err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	if err := h.services.WorkoutList.UpdateWorkout(id, input); err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, statusResponse{
// 		Status: "ok",
// 	})
// }

func (h *Handler) deleteNutrition(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.NutritionList.DeleteNutrition(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
