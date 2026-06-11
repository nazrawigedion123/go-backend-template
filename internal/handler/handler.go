package handler

import "github.com/gin-gonic/gin"

type SampleHandler interface {

	// Create handles POST /api/samples
	Create(c *gin.Context)

	GetAll(c *gin.Context)
}
