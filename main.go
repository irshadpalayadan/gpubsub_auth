package main

import (
	"github.com/gin-gonic/gin"
	auth "github.com/irshadpalayadan/gpubsub_auth/module/auth"
)

func getServerStatus(ctx *gin.Context) {
	ctx.JSON(200, "auth server is healthy")
}

func main() {
	auth.InitializeUser()
	router := gin.Default()

	router.GET("/status", getServerStatus)
	router.POST("/login", auth.SignIn)
	router.POST("/signup", auth.SignUp)

	router.Run(":8001")
}
