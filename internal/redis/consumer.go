package redis

import (
	"encoding/json"
	"github.com/tanmaytushar/Judge0/internal/models"
	"github.com/redis/go-redis/v9"
)

func PopSubmission() (*models.Submission,error){
	res,err := rdb.LPop(ctx,"submissions").Result()
	if err == redis.Nil {
		return nil,nil
	}
	if err != nil{
		return nil,err
	}
	var sub models.Submission
	if err := json.Unmarshal([]byte(res),&sub); err != nil{
		return nil, err;
	}
	return &sub,nil
}