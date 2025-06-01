package worker

import (
	"log"
	"github.com/tanmaytushar/Judge0/internal/redis"
	"github.com/tanmaytushar/Judge0/internal/docker"
	"github.com/tanmaytushar/Judge0/internal/utils"
)

func StartWorker() {
	go func() {
		for {
			result, err := redis.PopSubmission()
			if err != nil {
				continue // Redis error
			}
			if result == nil {
				continue // No submission, wait for next
			}

			log.Println("Processing Submission:", result.ID)

			err = utils.PrepareSubmissionFiles(result.ID, result.Code, result.ProblemID)
			if err != nil {
				log.Println("Error preparing files:", err)
				continue
			}

			err = docker.RunDockerJudge(result.ID)
			if err != nil {
				log.Println("Docker execution error:", err)
				continue
			}

			match, err := utils.CompareOutput(result.ID)
			if err != nil {
				log.Println("Comparison error:", err)
				continue
			}

			status := "Accepted"
			if !match {
				status = "Wrong Answer"
			}

			log.Printf("Submission %s Result: %s\n", result.ID, status)
		}
	}()
}

