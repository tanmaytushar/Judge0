package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/tanmaytushar/Judge0/internal/api"
	"github.com/tanmaytushar/Judge0/internal/redis"
	"github.com/tanmaytushar/Judge0/internal/worker"
)

func main(){
	redis.InitRedis()
	worker.StartWorker()

	r := gin.Default()
	r.GET("/health", func(c *gin.Context){
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.POST("/submit", api.Submit)
	
	log.Println("Running on port 8080")
	r.Run(":8080")
}
