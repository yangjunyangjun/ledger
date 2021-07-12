package webbase

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseServer struct {
}



type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (*BaseServer) Success(c *gin.Context, data interface{}) {
	rep := Response{
		Code: Success,
		Msg:  "ok",
		Data: data,
	}
	c.JSON(http.StatusOK, rep)
}

func (*BaseServer) Error(c *gin.Context, code int, msg string) {
	rep := Response{
		Code: code,
		Msg:  msg,
	}
	c.JSON(http.StatusOK, rep)
}
