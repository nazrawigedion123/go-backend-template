package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/nazrawigedion123/go-backend-template/platform/logger"
)

type Route struct {
	Method     string
	Path       string
	Handler    gin.HandlerFunc
	Middleware []gin.HandlerFunc
}

func RegisterRoute(grg *gin.RouterGroup, routes []Route, log logger.Logger) {
	for _, route := range routes {
		var handler []gin.HandlerFunc
		handler = append(handler, route.Middleware...)
		handler = append(handler, route.Handler)
		grg.Handle(route.Method, route.Path, handler...)
	}
}
