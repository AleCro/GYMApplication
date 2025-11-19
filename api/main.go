package main

import (
	Db "Svelgok-API/Database"
	"Svelgok-API/Environment"
	"Svelgok-API/Routes"

	"github.com/gin-gonic/gin"
)

var (
	API_VERSION uint8 = 1
)

func main() {
	_, err := Db.Connect(Environment.DATABASE_URL)
	if err != nil {
		panic(err)
	}

	err = Db.Connection.Initialize()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", Environment.API_CORS_ORIGIN)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", Environment.API_CORS_CREDENTIALS)
		c.Writer.Header().Set("Access-Control-Allow-Headers", Environment.API_CORS_HEADERS)
		c.Writer.Header().Set("Access-Control-Allow-Methods", Environment.API_CORS_METHODS)
		c.Next()
	})

	r.Use(Routes.SoftSession())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": API_VERSION,
		})
	})

	//      ____                _
	//     / ___|  ___  ___ ___(_) ___  _ __
	//     \___ \ / _ \/ __/ __| |/ _ \| '_ \
	//      ___) |  __/\__ \__ \ | (_) | | | |
	//     |____/ \___||___/___/_|\___/|_| |_|

	r.POST("/session", Routes.CreateJWTSession)
	r.PATCH("/session", Routes.RequireSessionValidateEx(), Routes.RefreshJWTSession)
	r.DELETE("/session", Routes.RemoveSession)
	r.OPTIONS("/session", Routes.OPTIONS)

	r.GET("/user", Routes.RequireSessionValidate(), Routes.GetSelf)
	r.POST("/user/change-password", Routes.RequireSessionValidate(), Routes.PasswordChange)
	r.OPTIONS("/user/change-password", Routes.OPTIONS)
	r.OPTIONS("/user", Routes.OPTIONS)

	r.GET("/users", Routes.RequireSessionValidate(), Routes.EnumerateUsers)
	r.POST("/users", Routes.CreateUser)
	r.OPTIONS("/users", Routes.OPTIONS)

	r.GET("/users/:id", Routes.RequireSessionValidate(), Routes.GetUsers)
	r.PUT("/users/:id", Routes.RequireSessionValidate(), Routes.EditUsers)
	r.OPTIONS("/users/:id", Routes.OPTIONS)

	// Notes
	r.GET("/notes", Routes.RequireSessionValidate(), Routes.GetNotes)
	r.POST("/notes", Routes.RequireSessionValidate(), Routes.CreateNote)
	r.PUT("/notes/:id", Routes.RequireSessionValidate(), Routes.UpdateNote)
	r.DELETE("/notes/:id", Routes.RequireSessionValidate(), Routes.DeleteNote)
	r.OPTIONS("/notes", Routes.OPTIONS)
	r.OPTIONS("/notes/:id", Routes.OPTIONS)

	// Events
	r.GET("/events", Routes.RequireSessionValidate(), Routes.GetEvents)
	r.POST("/events", Routes.RequireSessionValidate(), Routes.CreateEvent)
	r.PUT("/events/:id", Routes.RequireSessionValidate(), Routes.UpdateEvent)
	r.DELETE("/events/:id", Routes.RequireSessionValidate(), Routes.DeleteEvent)
	r.OPTIONS("/events", Routes.OPTIONS)
	r.OPTIONS("/events/:id", Routes.OPTIONS)

	// Progress
	r.GET("/progress", Routes.RequireSessionValidate(), Routes.GetProgress)
	r.POST("/progress", Routes.RequireSessionValidate(), Routes.CreateProgress)
	r.PUT("/progress/:id", Routes.RequireSessionValidate(), Routes.UpdateProgress)
	r.DELETE("/progress/:id", Routes.RequireSessionValidate(), Routes.DeleteProgress)
	r.OPTIONS("/progress", Routes.OPTIONS)
	r.OPTIONS("/progress/:id", Routes.OPTIONS)

	// Static files
	r.Static("/static", "./static")

	r.Run(":8081")
}
