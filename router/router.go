package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	chathandler "guizizhan/controller/chat"
	"guizizhan/controller/note"
	"guizizhan/controller/recruit"
	"guizizhan/controller/treasurehunting"
	"guizizhan/controller/user_behave"
	qiniu2 "guizizhan/pkg/qiniu"
	"guizizhan/pkg/token"
	"net/http"
)

func RouterInit(db *gorm.DB) *gin.Engine {
	a := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	a.Use(sessions.Sessions("mysession", store))

	// QiNiuTokenGroup 提供与 QiNiuToken 相关的端点。
	QiNiuTokenGroup := a.Group("/qiniutoken")
	{

		QiNiuTokenGroup.GET("", func(c *gin.Context) {
			qiniu2.QiniuInit()
			qnToken := qiniu2.GetQNToken()
			c.JSON(http.StatusOK, gin.H{
				"qnToken": qnToken,
			})
		})
	}

	// LoginGroup 提供与用户登录相关的端点。
	LoginGroup := a.Group("/api/login")
	{

		LoginGroup.POST("", func(c *gin.Context) {
			user_behave.Login(c, db)
		})
	}

	UserGroup := a.Group("/api/user") //.Use(middleware.Auth())
	{

		UserGroup.GET("/detail", token.JWTAuthMiddleware(), func(c *gin.Context) {
			user_behave.GetUserDetails(c, db)
		}) // 获取用户信息

		UserGroup.POST("/update", token.JWTAuthMiddleware(), func(c *gin.Context) {
			user_behave.UpdateUserInfo(c, db)
		}) // 更新用户信息

		UserGroup.GET("/avatar", token.JWTAuthMiddleware(), func(c *gin.Context) {
			user_behave.UpdateUserAvatar(c, db)
		}) // 更新用户头像

		UserGroup.GET("/achievement/get", func(c *gin.Context) {
			user_behave.GetAchievements(c, db)
		})
		UserGroup.GET("/achievement/update", func(c *gin.Context) {
			user_behave.UpdateAchievement(c, db)
		})
	}

	ContactGroup := a.Group("/api/other")
	{

		ContactGroup.GET("/follow", token.JWTAuthMiddleware(), func(c *gin.Context) {
			user_behave.Follow(c, db)
		})
	}

	UserPostGroup := a.Group("/api/post")
	{

		UserPostGroup.POST("/postnote", token.JWTAuthMiddleware(), func(c *gin.Context) {
			note.PostNote(c, db)
		}) // 发布记录

		UserPostGroup.POST("/post_treasure_hunting", token.JWTAuthMiddleware(), func(c *gin.Context) {
			treasurehunting.PostTreasureHunting(c, db)
		}) // 发布寻宝活动

		UserPostGroup.POST("/post_recruit_activity", token.JWTAuthMiddleware(), func(c *gin.Context) {
			recruit.PostRecruit(c, db)
		}) // 发布招募活动
	}

	GetActivityGroup := a.Group("/api/getactivity")
	{
		// @Summary 获取所有记录
		// @Description 获取所有发布的记录的详细信息。
		GetActivityGroup.GET("/allpostnote", func(c *gin.Context) {
			note.GetAllPostNote(c, db)
		})

		//GetActivityGroup.GET("/thepostnote", token.JWTAuthMiddleware(), func(c *gin.Context) {
		//	note.GetThePostNote(c, db)
		//})

		GetActivityGroup.GET("/alltreasurehunting", func(c *gin.Context) {
			treasurehunting.GetAllTreasureHuntings(c, db)
		})

		//GetActivityGroup.GET("/thetreasurehunting", token.JWTAuthMiddleware(), func(c *gin.Context) {
		//	treasurehunting.GetTheTreasureHunting(c, db)
		//})

		GetActivityGroup.GET("/allrecruit", func(c *gin.Context) {
			recruit.GetAllRecruits(c, db)
		})

		//GetActivityGroup.GET("/therecruit", token.JWTAuthMiddleware(), func(c *gin.Context) {
		//	recruit.GetTheRecruit(c, db)
		//})
	}
	ChatGroup := a.Group("/api/talk")
	{
		ChatGroup.GET("", token.JWTAuthMiddleware(), func(c *gin.Context) {
			chathandler.WsHandler(c, db)
		})
	}
	a.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})

	return a
}
