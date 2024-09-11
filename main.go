package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kalyanKumarPokkula/Go-jwt/controllers"
	"github.com/kalyanKumarPokkula/Go-jwt/initializers"
	"github.com/kalyanKumarPokkula/Go-jwt/middlewares"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.Connect()
	initializers.SyncDatabase()
	initializers.GoogleConfig()
}

func Hello(c *gin.Context){
	c.JSON(200, gin.H{
        "message": "Hello World",
    })
}

func main(){
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Specify origins you want to allow
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// r.Use(func(c *gin.Context) {
    //     c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
    //     c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 	// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    //     if c.Request.Method == "OPTIONS" {
    //         c.AbortWithStatus(http.StatusNoContent)
    //         return
    //     }
    //     c.Next()
    // })


    // Serve static files without conflicting paths
    // r.Static("/assets", "./static")

	r.GET("/api/someendpoint", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, this is a sample endpoint!",
        })
    })

	r.POST("/api/signup" , controllers.Signup)
	r.POST("/api/signin" , controllers.Login)
	r.GET("/api/google_login" , controllers.GoogleLogin);
	r.GET("/api/google_callback", controllers.GoogleCallback);
	r.GET("/api/validate" ,middlewares.AuthenticateJwt , controllers.Vaildate)

	r.Run(":3002")
}

