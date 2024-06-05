package routes

import (
	"github.com/gin-gonic/gin"
)

// InitializeRoutes sets up all the routes
func InitializeRoutes(r *gin.Engine) {
	InitializeUserRoutes(r)
}
