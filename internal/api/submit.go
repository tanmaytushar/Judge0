package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tanmaytushar/Judge0/internal/models"
	"github.com/tanmaytushar/Judge0/internal/redis"
)

func Submit(c *gin.Context){
	var sub models.Submission
	err := c.ShouldBindJSON(&sub);
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub.ID = uuid.New().String()
	sub.Status = "pending"
	sub.Verdict = "waiting"

	// Push to redis queue
	redis.EnqueueSubmission(sub);

	c.JSON(http.StatusOK,gin.H{
		"message" : "Submission recieved",
		"id": sub.ID,
	})

}