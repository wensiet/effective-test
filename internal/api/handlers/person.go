package handlers

import (
	"effective-test/internal/api/schemas"
	"effective-test/internal/models"
	"effective-test/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CreatePerson godoc
// @Summary Create new person
// @Description Allows to create person and predict several characteristics
// @Tags Persons
// @Accept json
// @Param person body schemas.CreatePerson true "Person's data"
// @Produce json
// @Success 200 {object} schemas.ResponsePerson
// @Router /person [post]
func CreatePerson(context *gin.Context) {
	var body schemas.CreatePerson
	err := context.Bind(&body)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	person, err := service.CreatePerson(body)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(201, person.AsResponse())
}

// DeletePerson godoc
// @Summary Delete person
// @Description Allows to delete person using person's ID
// @Tags Persons
// @Accept json
// @Param id path string true "Person's ID"
// @Produce json
// @Success 200
// @Router /person/{id} [delete]
func DeletePerson(context *gin.Context) {
	personID := context.Param("id")

	err := service.DeletePersonByID(personID)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "Deleted"})
}

// UpdatePerson godoc
// @Summary Update person
// @Description Allows to update person using person's ID and update data
// @Tags Persons
// @Accept json
// @Param id path string true "Person's ID"
// @Param person body schemas.UpdatePerson true "Update data"
// @Produce json
// @Success 200
// @Router /person/{id} [put]
func UpdatePerson(context *gin.Context) {
	personID := context.Param("id")

	var body schemas.UpdatePerson
	err := context.Bind(&body)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = service.UpdatePerson(personID, body)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "Updated"})
}

// GetPerson godoc
// @Summary Get one person
// @Description Allows to get person using person's ID
// @Tags Persons
// @Accept json
// @Param id path string true "Person's ID"
// @Produce json
// @Success 200
// @Router /person/{id} [get]
func GetPerson(context *gin.Context) {
	personID := context.Param("id")

	var person models.Person
	err := person.Find(personID)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, person.AsResponse())
}

// GetPersonsWithPaging godoc
// @Summary Get 10 person
// @Description Allows to get 10 person using paging, each page has 10 elements
// @Tags Persons
// @Accept json
// @Param num query string true "Page number"
// @Produce json
// @Success 200
// @Router /person/page [get]
func GetPersonsWithPaging(context *gin.Context) {
	pageNum, err := strconv.Atoi(context.Request.URL.Query().Get("num"))
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	results, err := service.GetPersonsWithPaging(pageNum)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, results)
}
