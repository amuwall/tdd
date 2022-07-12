package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"tdd/dao"
	"tdd/web/response"
)

func GetUsers(c *gin.Context) {
	type QueryParams struct {
		Page     uint32 `form:"page" binding:"required,gt=0"`
		PageSize uint32 `form:"page_size" binding:"required,gt=0"`
	}

	var queryParams QueryParams
	err := c.ShouldBindQuery(&queryParams)
	if err != nil {
		log.Printf("GetUsers ShouldBindQuery error = %v\n", err)
		response.Error(c, response.ErrorCodeInvalidParams, "invalid params")
		return
	}

	users, err := dao.GetUsers(dao.GetDB(), queryParams.Page, queryParams.PageSize)
	if err != nil {
		log.Printf("dao.GetUsers error = %v\n", err)
		response.Error(c, response.ErrorCodeDatabaseError, "get users from database error")
		return
	}

	response.Success(c, gin.H{"users": users})
}
