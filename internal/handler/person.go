package handler

import (
	"em-test/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type userUpdate struct {
	Id     int             `json:"id"`
	Person models.PersonBD `json:"person"`
}

type inputId struct {
	Id int `json:"id"`
}

func (h *Handler) AddPerson(c *gin.Context) {
	logrus.Info("Post(AddPerson) request received")
	var input models.PersonInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Person.AddPerson(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Info("Post(AddPerson) request complete")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
func (h *Handler) GetPersons(c *gin.Context) {
	logrus.Info("Get (person) request received")
	var input models.UserGetList
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(input)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	persons, err := h.services.Person.GetPersons(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Info("Get (person) request complete")
	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}
func (h *Handler) EditPerson(c *gin.Context) {
	logrus.Info("Patch (person) request received")
	var input models.PersonBD
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Person.EditPerson(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Info("Patch (person) request complete")

	c.JSON(http.StatusOK, gin.H{
		"message": "Done",
	})
}
func (h *Handler) DeletePerson(c *gin.Context) {
	logrus.Info("Delete (person) request received")
	var input inputId
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Person.DeletePerson(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Info("Delete (person) request complete")

	c.JSON(http.StatusOK, gin.H{
		"message": "Done",
	})
}
