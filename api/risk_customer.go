package api

import (
	"backend/model"
	"backend/response"
	"github.com/gin-gonic/gin"
)

func GetBanCustomerList(c *gin.Context) {
	customerList, err := model.GetCustomerList()
	if err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return
	}
	response.Ok().Data(customerList).Send(c)
}
