package main

import (
	// "context"
	// "encoding/json"
	"fmt"
	"time"

	"github.com/tanmaytushar/Judge0/internal/redis"
)

func main(){
	redis.InitRedis()
	fmt.Println("Worker started")

	for{
		result, err := redis.PopSubmission()
		if err != nil{
			fmt.Println("Errror Popping Submissin: ",err)
			// wait for items to get into redis queue
			time.Sleep(time.Second)
			continue
		}
		if result != nil{
			fmt.Println("Processing Submission:",result.ID)
			// TODO : docker and judge logic here
		}
	}
}




