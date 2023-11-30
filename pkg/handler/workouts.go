package handler

import (
	"net/http"
	"sportnotes"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getAllWorkoutsResponse struct {
	Data []sportnotes.WorkoutOutputAll `json:"data"`
}

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

	id, err := h.services.WorkoutList.CreateWorkout(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newWorkoutResponse{
		Id: id,
	})
}

func (h *Handler) getAllWorkouts(c *gin.Context) {
	workouts, err := h.services.WorkoutList.GetAllWorkouts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllWorkoutsResponse{
		Data: workouts,
	})
}

func (h *Handler) getWorkoutById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	workout, err := h.services.WorkoutList.GetWorkoutById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, workout)
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
