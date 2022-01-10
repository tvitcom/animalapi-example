package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetInt64IdFromReqContext(c *gin.Context) int64 {
	IdStr := c.Param("id")
	id, err := strconv.ParseInt(IdStr, 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return id
}

func GetLimitOffset(c *gin.Context) (int64, int64, int64) {
	pageParam, _ := c.GetQuery("page")
	page, _ := strconv.ParseInt(pageParam, 10, 64)
	if page == 0 {
		page = 1
	}

	limit := int64(10000)
	offset := limit * (page - 1)

	return limit, offset, page
}
