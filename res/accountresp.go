package res

import (
	"devops/myerr"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountResp struct {
	AccountName string `json:"accountName"`
}

type ListResponse struct {
	TotalCount  uint64         `json:"totalCount"`
	AccountList []*AccountResp `json:"accountList"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := myerr.DecodeErr(err)

	//http.StatusOK这个值是200
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
