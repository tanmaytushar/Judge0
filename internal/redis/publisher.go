package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/tanmaytushar/Judge0/internal/models"
)

var ctx = context.Background()
var rdb *redis.Client

func InitRedis(){
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func EnqueueSubmission(sub models.Submission) error {
	data, _ := json.Marshal(sub)
	return rdb.RPush(ctx,"submissions",data).Err()
}