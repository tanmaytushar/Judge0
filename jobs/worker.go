package main

import (
	// "context"
	// "encoding/json"
	"fmt"
	"time"
	"github.com/tanmaytushar/Judge0/internal/redis"
	"github.com/tanmaytushar/Judge0/internal/utils"
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
			// TODO : Judge Logic
			err := utils.PrepareSubmissionFiles(result.ID, result.Code, result.ProblemID)
			if err != nil {
				fmt.Println("Failed to prepare submission files:", err)
				continue
			}
			fmt.Println("Submission files prepared for:", result.ID)
			// TODO : docker here
		}
	}
}




