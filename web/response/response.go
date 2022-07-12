package response

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Body struct {
	Code ErrorCode   `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (body *Body) String() string {
	data, _ := json.Marshal(body)
	return string(data)
}

func Success(c *gin.Context, data interface{}) {
	body := Body{
		Code: ErrorCodeSuccess,
		Msg:  "",
		Data: data,
	}

	c.JSON(http.StatusOK, body)
}

func Error(c *gin.Context, code ErrorCode, msg string) {
	body := Body{
		Code: code,
		Msg:  msg,
		Data: nil,
	}

	c.JSON(http.StatusOK, body)
}

func ErrorWithData(c *gin.Context, code ErrorCode, msg string, data interface{}) {
	body := Body{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	c.JSON(http.StatusOK, body)
}
