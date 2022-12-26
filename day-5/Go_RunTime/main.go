package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type job_struct struct {
	Id   int `json:"id,string,omitempty"`
	Name string
}

var jobs_struct = []job_struct{
	{Id: 342, Name: "build-app"},
	{Id: 123, Name: "build-server"},
	{Id: 78, Name: "build-shared-library"},
}

type result struct {
	id     int
	status string
}

func execute(job job_struct) result {
	return result{id: job.Id, status: "SUCCESS"}
}

func setupRouter() *gin.Engine {

	router := gin.Default()

	router.POST("/add-job", func(c *gin.Context) {

		Id := c.PostForm("Id")
		Id1, err := strconv.Atoi(Id)
		if err != nil {
			fmt.Println(err)
		}
		Name := c.PostForm("Name")

		data := append(jobs_struct, job_struct{Id: Id1, Name: Name})

		ch := make(chan result, len(data))
		for _, job := range data {
			go func(job job_struct) {
				ch <- execute(job)
			}(job)
		}

		results := []result{}

		for range data {
			results = append(results, <-ch)
		}

		c.JSON(200, gin.H{
			"message":        "OK",
			"results_length": len(results),
		})
	})

	return router
}

func main() {

	router := setupRouter()
	router.Run(":8080")
}
