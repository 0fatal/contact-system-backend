package api

import (
	"backend/middleware"
	"backend/response"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Data map[string]interface{}

func Init() *gin.Engine {
	_r := gin.Default()

	_r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     []string{"", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	})) // 跨域

	_r.Use(middleware.Auth()) // 身份认证

	r := _r.Group("api")
	r.POST("/login", Login)
	r.POST("/upload", Upload)
	r.GET("/customer", GetBanCustomerList)

	user := r.Group("user")
	{
		user.GET("/info", GetInfo)
	}

	identity := r.Group("identity")

	{
		identity.POST("/", CreateNewIdentify)
		identity.POST("/check", CheckIdentityApply)
		identity.GET("/list", GetIdentifyApplyList)
		identity.GET("/:id", GetIdentifyDetail)
	}

	refresh := r.Group("refresh")
	{
		refresh.POST("/", CreateNewRefresh)
		refresh.POST("/check", CheckRefreshApply)
		refresh.GET("/list", GetRefreshApplyList)
		refresh.GET("/:id", GetRefreshDetail)
	}

	reason := r.Group("reason")
	{
		reason.GET("/list", GetReasonList)
		reason.GET("/list/all", GetAllReasonList)
		reason.GET("/:id", GetReasonDetail)
		reason.POST("/:id", EnableReason)
	}

	return _r
}

func HandleDTOVerifyError(err error, c *gin.Context) bool {
	if err != nil {
		response.Fail().Msg(err.Error()).Send(c)
		return true
	}
	return false
}
