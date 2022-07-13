package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"tdd/dao"
	"tdd/model"
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

func SearchUsers(c *gin.Context) {
	type BodyParams struct {
		IDs []uint32 `json:"ids" binding:"required,dive,gt=0"`
	}

	var bodyParams BodyParams
	err := c.ShouldBindJSON(&bodyParams)
	if err != nil {
		log.Printf("SearchUsers ShouldBindJson error = %v\n", err)
		response.Error(c, response.ErrorCodeInvalidParams, "invalid params")
		return
	}

	users := []*model.User{}
	// 此处作为 gomonkey.ApplyFuncSeq 的 Demo，所以采用 for 循环，依次查询 user_id
	// 实际业务场景中，应直接使用数据库的 IN 查询功能
	for _, userID := range bodyParams.IDs {
		user, err := dao.GetUserByID(dao.GetDB(), userID)
		if err != nil {
			log.Printf("SearchUsers GetUserByID error = %v\n", err)
			response.Error(c, response.ErrorCodeDatabaseError, "search user by id from database error")
			return
		}

		users = append(users, user)
	}

	response.Success(c, gin.H{"users": users})
}
