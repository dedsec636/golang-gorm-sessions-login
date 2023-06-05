package main

import (
	"ginbasic/handlers"
	"ginbasic/helper"
	"ginbasic/middleware"
	"ginbasic/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	//load db using env
	loadEnv()
	LoadDatabase()

	r.GET("/signup", handlers.GetSignupHandler)
	r.POST("/signup", handlers.PostSignupHandler)

	r.GET("/login", handlers.GetLoginHandler)
	r.POST("/login", handlers.PostLoginHandler)

	r.GET("/home", middleware.SessionMiddleware(), handlers.ProtectedHomePageHandler)
	r.GET("/logout", handlers.LogoutHandler)

	r.GET("/create", middleware.SessionMiddleware(), handlers.GetCreateNewHandler)
	r.POST("/createAuth", middleware.SessionMiddleware(), handlers.CreateAuth)
	r.POST("/createBook", middleware.SessionMiddleware(), handlers.CreateBook)
	err := r.Run(":8000")
	helper.ErrorPanic(err)
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		helper.ErrorPanic(err)
	}
}

func LoadDatabase() {
	helper.ConnectDb()
	//helper.Database.AutoMigrate(&model.User{})
	helper.Database.AutoMigrate(&model.Author{})
	helper.Database.AutoMigrate(&model.Book{})
}
