package handler

import (
	"em-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddPerson(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
func (h *Handler) CheckPerson(c *gin.Context) {

}
func (h *Handler) EditPerson(c *gin.Context) {
	var input userUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Person.EditPerson(input.Id, input.Person)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Done",
	})
}
func (h *Handler) DeletePerson(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Done",
	})
}

type inputId struct {
	Id int `json:"id"`
}
type userUpdate struct {
	Id     int             `json:"id"`
	Person models.PersonBD `json:"person"`
}
