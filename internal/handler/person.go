package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type personInput struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

func (h *Handler) AddPerson(c *gin.Context) {
	var input personInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Person.AddPerson()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

}
func (h *Handler) CheckPerson(c *gin.Context) {

}
func (h *Handler) EditPerson(c *gin.Context) {

}
func (h *Handler) DeletePerson(c *gin.Context) {

}
