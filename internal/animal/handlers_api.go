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

	// c.JSON(http.StatusOK, "animal/index.tmpl", m)

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
		c.JSON(http.StatusNotFound,gin.H{"ok": false, "data": ""})
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
		c.JSON(http.StatusNotFound, gin.H{"ok": false, "data": ""})
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
	bindErr := c.ShouldBind(&animal)
	if bindErr != nil {
		animal.Error = bindErr
		c.HTML(http.StatusOK, "animal/edit.tmpl", animal)
		return
	}

	foundAnimal:= model.FindById(id)
	// Check if resource exist
	if foundAnimal.Id == 0 {
		c.HTML(http.StatusNotFound, "common/not_found.tmpl", gin.H{})
	}

	// Updating data
	animal, updateErr := model.Put(foundAnimal.Id, animal)
	if updateErr != nil {
		c.HTML(http.StatusInternalServerError, "common/internal_error.tmpl", gin.H{})
		util.PanicError(updateErr)
	} else {
		c.Redirect(http.StatusFound, "/animal/show/"+strconv.FormatInt(id, 10))
	}
}

func ApiDeleteHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	animal:= model.FindById(id)

	// Check if resource exist
	if animal.Id == 0 {
		c.HTML(http.StatusNotFound, "common/not_found.tmpl", gin.H{})
		return
	}

	err := model.Delete(animal)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "common/internal_error.tmpl", gin.H{})
		return
	} else {
		c.Redirect(http.StatusFound, "/animal/")
	}
}
