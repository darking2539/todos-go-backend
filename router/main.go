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
	authRouter.GET("/login", handler.LoginHandler)
	authRouter.GET("/logout", handler.LogoutHandler)
	authRouter.GET("/getuser", middleware.JWTAuthUser, handler.GetUserHandler)
	authRouter.GET("/callback", handler.CallbackHandler)
	authRouter.POST("/register", handler.RegisterHandler)
	authRouter.POST("/changepassword", middleware.JWTAuthUser, handler.ChangePasswordHandler)

	//todos
	todosRouter := router.Group("/todos")
	todosRouter.Use(middleware.JWTGetUserMiddleware())
	todosRouter.POST("/submit", handler.SubmitTodosHandler)
	todosRouter.GET("/list", handler.GetListTodosHandler)
	todosRouter.DELETE("/:id", handler.DeleteTodosHandler)

}
