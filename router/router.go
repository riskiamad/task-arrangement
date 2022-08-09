package router

import "github.com/gin-gonic/gin"

type RouteHandlers interface {
	URLMapping(r *gin.RouterGroup)
}

var handlers = map[string]RouteHandlers{}

func Router() *gin.Engine {
	router := gin.Default()
	v := router.Group("v1/")
	if len(handlers) > 0 {
		for r, h := range handlers {
			h.URLMapping(v.Group(r))
		}
	}

	return router
}
