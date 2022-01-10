package animal

import (
	"github.com/gin-gonic/gin"
	"github.com/tvitcom/animalapi-example/internal/animal/model"
	"github.com/tvitcom/animalapi-example/pkg/util"
	"net/http"
	"strconv"
)

type AnimalItem struct {
	Kind string
	Name string
	Dob int
	Owner string
}
/*
{
	"ok":true,
	"data":""
}
*/


func ApiIndexHandler(c *gin.Context) {
	limit, offset, page := util.GetLimitOffset(c)
	animals := model.IndexWithPage(limit, offset)
	count := model.Count()
	pagination := util.ProcessPagination("animal", count, page, limit)

	m := make(map[string]interface{})
	m["animals"] = animals
	m["pagination"] = pagination

	c.JSON(http.StatusOK, gin.H{
		"ok":true,
		"data": m,
	})
}

func ApiShowHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	animal := model.FindById(id)

	// Check if resource exist
	if animal.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"ok": false, "data": "No data found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":true,
		"data":animal,
	})
}

func ApiNewHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok":true,
		"data":`"json-template:{"kind":"required","name":"required","dob":"gte=19991231","owner":"required"}"`})
}

func ApiCreateHandler(c *gin.Context) {
	var animal model.Animal
	if err := c.ShouldBindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "data": err.Error()})
		return
	}

	// Inserting data
	id, err := model.Create(animal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "data": err.Error()})
		util.PanicError(err)
	}

	// c.Redirect(http.StatusFound, "/animal/show/"+strconv.FormatInt(id, 10))
	animal = model.FindById(id)
	// Check if resource exist
	if animal.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"ok": false, "data": "No data found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok":true,
		"data":"/animal/show/" + strconv.FormatInt(animal.Id, 10),
	})
}

func ApiUpdateHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	var animal model.Animal

	// App level validation
	err := c.ShouldBindJSON(&animal)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"ok": false, "data": err.Error()})
		return
	}

	foundAnimal:= model.FindById(id)
	// Check if resource exist
	if foundAnimal.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"ok": false, "data": "Absent data"})
	}

	// Updating data
	animal, updateErr := model.Put(foundAnimal.Id, animal)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "data": updateErr.Error()})
		util.PanicError(updateErr)
	} 
	c.JSON(http.StatusOK, gin.H{
		"ok":true,
		"data":"/animal/show/" + strconv.FormatInt(animal.Id, 10),
	})
}

func ApiDeleteHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	animal:= model.FindById(id)

	// Check if resource exist
	if animal.Id == 0 {
		c.JSON(http.StatusNotFound,gin.H{"ok": false, "data": "already absent data"})
		return
	}

	err := model.Delete(animal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok":true,
		"data":"",
	})
}
