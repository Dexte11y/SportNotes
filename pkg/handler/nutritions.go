package handler

import (
	"net/http"
	"sportnotes"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type getWorkoutsByParamResponse struct {
// 	Data []sportnotes.WorkoutOutputByParam `json:"data"`
// }

func (h *Handler) createNutrition(c *gin.Context) {
	// _, err := getDoctorId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	var input sportnotes.Nutrition
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.NutritionList.CreateNutrition(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newNutritionResponse{
		Id: id,
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
// 	// _, err := getDoctorId(c)
// 	// if err != nil {
// 	// 	return
// 	// }

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
	// _, err := getDoctorId(c)
	// if err != nil {
	// 	return
	// }

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
