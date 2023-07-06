package router

import (
	ctl "todos-go-backend/controller"
	"todos-go-backend/db"
	"todos-go-backend/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {

	router.GET("/healthz", db.Healthz)

	//auth
	authRouter := router.Group("/auth")
	authRouter.POST("/login", ctl.LoginCtl)
	authRouter.POST("/register", ctl.RegisterCtl)
	authRouter.POST("/changepassword", middleware.JWTAuthUser, ctl.ChangePasswordCtl)

	//todos
	todosRouter := router.Group("/todos")
	todosRouter.Use(middleware.JWTGetUserMiddleware())
	todosRouter.POST("/submit", ctl.SubmitTodosCtl)
	todosRouter.GET("/list", ctl.GetListTodosCtl)
	todosRouter.DELETE("/:id", ctl.DeleteTodosCtl)

}
