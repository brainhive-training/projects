package main

import (
	"github.com/gin-gonic/gin"

	"day5/project-api/project"
)

func main() {
	r := gin.Default()

	project.RegisterHandlers(r, &project.Service{})

	r.Run(":8080")
}
