package router

import (
	handler "todos-go-backend/handler"
	"todos-go-backend/db"
	"todos-go-backend/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {

	router.GET("/healthz", db.Healthz)

	//auth
	authRouter := router.Group("/auth")
	authRouter.POST("/login", handler.LoginCtl)
	authRouter.POST("/register", handler.RegisterCtl)
	authRouter.POST("/changepassword", middleware.JWTAuthUser, handler.ChangePasswordCtl)

	//todos
	todosRouter := router.Group("/todos")
	todosRouter.Use(middleware.JWTGetUserMiddleware())
	todosRouter.POST("/submit", handler.SubmitTodosCtl)
	todosRouter.GET("/list", handler.GetListTodosCtl)
	todosRouter.DELETE("/:id", handler.DeleteTodosCtl)

}
