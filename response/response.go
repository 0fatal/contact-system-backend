package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type R struct {
	code int
	msg  string
	data interface{}
}

func New(code int, msg string, data interface{}) *R {
	return &R{
		code: code,
		msg:  msg,
		data: data,
	}
}

func Ok() *R {
	return &R{
		code: 0,
		msg:  "success",
		data: map[string]interface{}{},
	}
}

func Fail() *R {
	return &R{
		code: -1,
		msg:  "error",
		data: nil,
	}
}

func (r *R) Code(code int) *R {
	r.code = code
	return r
}

func (r *R) Msg(msg string) *R {
	r.msg = msg
	return r
}

func (r *R) Data(data interface{}) *R {
	r.data = data
	return r
}

func (r *R) Send(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": r.code,
		"msg":  r.msg,
		"data": r.data,
	})
}
