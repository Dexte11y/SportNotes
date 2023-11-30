package handler

import (
	"fmt"
	"net/http"
	"sportnotes"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getAllTrainingsResponse struct {
	Data []sportnotes.Training `json:"data"`
}

func (h *Handler) createTraining(c *gin.Context) {
	// _, err := getDoctorId(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	var input sportnotes.Training
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(input)

	id, err := h.services.TrainingList.CreateTraining(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newUserResponse{
		Id: id,
	})
}

func (h *Handler) getAllTrainings(c *gin.Context) {
	trainings, err := h.services.TrainingList.GetAllTrainings()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTrainingsResponse{
		Data: trainings,
	})
}

func (h *Handler) getTrainingById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := h.services.TrainingList.GetTrainingById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) updateTraining(c *gin.Context) {
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

func (h *Handler) deleteTraining(c *gin.Context) {
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
