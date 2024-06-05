// helpers/register_routes.go
package helpers

import (
	"github.com/gin-gonic/gin"
)

// RouteConfig defines the structure of a route configuration
type RouteConfig struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

// RegisterRoutes registers a list of routes to the given router
func RegisterRoutes(r *gin.Engine, routes []RouteConfig) {
	for _, route := range routes {
		registerRoute(r, route.Method, route.Path, route.HandlerFunc)
	}
}

// registerRoute registers a single route to the given router
func registerRoute(r *gin.Engine, method, path string, handlerFunc gin.HandlerFunc) {
	switch method {
	case "GET":
		r.GET(path, handlerFunc)
	case "POST":
		r.POST(path, handlerFunc)
	case "PUT":
		r.PUT(path, handlerFunc)
	case "DELETE":
		r.DELETE(path, handlerFunc)
	default:
		panic("Unsupported method: " + method)
	}
}
