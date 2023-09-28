package handler

import (
	"fmt"
	"net/http"
	"strconv"

	server "github.com/deuuus/bmstu-rsoi"
	"github.com/gin-gonic/gin"
)

func getPersonId(c *gin.Context) (int, error) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	return idInt, err
}

func (h *Handler) getPerson(c *gin.Context) {
	id, err := getPersonId(c)

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	person, err := h.services.Person.GetPersonById(id)

	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, person)
}

func (h *Handler) getAllPersons(c *gin.Context) {
	persons, err := h.services.Person.GetAllPersons()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, persons)
}

func (h *Handler) createPerson(c *gin.Context) {
	var input server.Person

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Person.CreatePerson(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Location", fmt.Sprintf("/api/v1/persons/%s", strconv.Itoa(id)))

	c.JSON(http.StatusCreated, "")
}

func (h *Handler) updatePerson(c *gin.Context) {
	id, err := getPersonId(c)

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input server.PersonUpdate
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	person, err := h.services.Person.UpdatePerson(id, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, person)
}

func (h *Handler) deletePerson(c *gin.Context) {
	id, err := getPersonId(c)

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Person.DeletePersonById(id)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, "")
}
