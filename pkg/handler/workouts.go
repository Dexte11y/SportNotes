package handler

import (
	"fmt"
	"net/http"
	"sportnotes"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type getWorkoutsByParamResponse struct {
// 	Data []sportnotes.WorkoutOutputByParam `json:"data"`
// }

func (h *Handler) createWorkout(c *gin.Context) {
	// _, err := getDoctorId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	var input sportnotes.Workout
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(input)

	id, err := h.services.WorkoutList.CreateWorkout(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newWorkoutResponse{
		Id: id,
	})
}

func (h *Handler) getWorkoutsByParam(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	input := c.DefaultQuery("interval", "none")
	if input == "" {
		input = "week"
	} //сделать обработку ошибки

	workouts, err := h.services.WorkoutList.GetWorkoutsByParam(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, workouts)
}

// func (h *Handler) updateWorkout(c *gin.Context) {
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

func (h *Handler) deleteWorkout(c *gin.Context) {
	// _, err := getDoctorId(c)
	// if err != nil {
	// 	return
	// }

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.WorkoutList.DeleteWorkout(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
