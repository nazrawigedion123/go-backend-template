package samplehandler

import (
	"net/http"

	"github.com/OnePulseOmni/pulse-wallet/internal/constant/dto"
	"github.com/OnePulseOmni/pulse-wallet/internal/constant/model/db"
	"github.com/OnePulseOmni/pulse-wallet/internal/constant/response"
	"github.com/OnePulseOmni/pulse-wallet/internal/handler"
	"github.com/OnePulseOmni/pulse-wallet/internal/module"
	"github.com/OnePulseOmni/pulse-wallet/platform/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type sampleHandler struct {
	sampleModule module.SampleModule
	logger       logger.Logger
}

// New creates a new sample handler instance
func New(logger logger.Logger, sampleModule module.SampleModule) handler.SampleHandler {
	return &sampleHandler{
		logger:       logger,
		sampleModule: sampleModule,
	}
}

// Create handles POST /api/samples
func (h *sampleHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	// Parse and validate request body using Gin binding
	var req dto.CreateSampleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error(ctx, "failed to bind request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
		return
	}

	// Prepare params
	params := db.CreateSampleParams{
		Name:  req.Name,
		Email: req.Email,
	}

	// Set optional fields
	if req.Age != 0 {
		params.Age = &req.Age

	}

	params.IsActive = &req.IsActive

	// Call module
	sample, err := h.sampleModule.Create(ctx, params)
	if err != nil {
		h.logger.Error(ctx, "failed to create sample", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Error:   "Failed to create sample",
			Message: err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, dto.CreateSampleResponse{
		Message: "Sample created successfully",
		Sample:  sample,
	})
}

// GetAll handles GET /api/samples
func (h *sampleHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	samples, err := h.sampleModule.GetAll(ctx)
	if err != nil {
		h.logger.Error(ctx, "failed to get all samples", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Error:   "Failed to fetch samples",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"samples": samples,
		"count":   len(samples),
	})
}

// package samplehandler

// import (
// 	"github.com/OnePulseOmni/pulse-wallet/internal/module"
// 	"github.com/OnePulseOmni/pulse-wallet/platform/logger"
// 	"github.com/gin-gonic/gin"
// )

// // RegisterRoutes registers sample routes
// func RegisterRoutes(r *gin.RouterGroup, logger logger.Logger, sampleModule module.SampleModule) {
// 	handler := New(logger, sampleModule)

// 	sampleGroup := r.Group("/samples")
// 	{
// 		sampleGroup.POST("/", handler.Create)  // POST /api/samples
// 		sampleGroup.GET("/", handler.GetAll)   // GET /api/samples
// 	}
// }
