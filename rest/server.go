package rest

import (
	"io/ioutil"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	addr string

	todoController TodoController
}

//NewHttpServer create server instance
func NewHttpServer() *HttpServer {
	return &HttpServer{
		addr: ":9001",
	}
}

func (server *HttpServer) Start() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	{
		r.GET("/api/task", server.todoController.GetTasks)
		r.POST("/api/task", server.todoController.AddTask)
	}

	err := r.Run(server.addr)
	if err != nil {
		panic(err)
	}
}
