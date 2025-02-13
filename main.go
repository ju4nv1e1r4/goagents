package main

import (
	"goagents/careerApp"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Role string `json:"role"`
	CV string `json:"cv"`
}


func analysis(c *gin.Context) {
	var content Input

	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	report := careerApp.Report(content.Role, content.CV)

	c.JSON(http.StatusOK, gin.H{"report": report})
}

func main()  {
	r := gin.Default()
	
	r.POST("/analysis", analysis)
	r.Run()
}