package animal

import (
	"github.com/gin-gonic/gin"
	"github.com/tvitcom/animalapi-example/internal/animal/model"
	"github.com/tvitcom/animalapi-example/pkg/util"
	"net/http"
	"strconv"
)

func ShowHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	item := model.FindById(id)

	// Check if resource exist
	if item.Id == 0 {
		c.HTML(http.StatusNotFound, "common/not_found.tmpl", gin.H{})
	} else {
		c.HTML(http.StatusOK, "animal/show.tmpl", item)
	}
}

func NewHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "animal/new.tmpl", gin.H{})
}

func CreateHandler(c *gin.Context) {
	var animal model.Animal

	// App level validation
	bindErr := c.ShouldBind(&animal)
	if bindErr != nil {
		animal.Error = bindErr
		c.HTML(http.StatusOK, "animal/new.tmpl", animal)
		return
	}

	// Inserting data
	id, insertErr := model.Create(animal)
	if insertErr != nil {
		c.HTML(http.StatusInternalServerError, "common/internal_error.tmpl", gin.H{})
		util.PanicError(insertErr)
	} else {
		c.Redirect(http.StatusFound, "/animal/show/"+strconv.FormatInt(id, 10))
	}
}

func EditHandler(c *gin.Context) {
	id := util.GetInt64IdFromReqContext(c)
	animal := model.FindById(id)

	// Check if resource exist
	if animal.Id == 0 {
		c.HTML(http.StatusNotFound, "common/not_found.tmpl", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "animal/edit.tmpl", animal)
}

func UpdateHandler(c *gin.Context) {
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

func IndexHandler(c *gin.Context) {
	limit, offset, page := util.GetLimitOffset(c)
	animals := model.IndexWithPage(limit, offset)
	count := model.Count()
	pagination := util.ProcessPagination("shopping-list", count, page, limit)

	m := make(map[string]interface{})
	m["animals"] = animals
	m["pagination"] = pagination

	c.HTML(http.StatusOK, "animal/index.tmpl", m)
}

func DeleteHandler(c *gin.Context) {
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
