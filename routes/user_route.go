// routes/user_routes.go
package routes

import (
	auth "go-crud/controllers/auth"
	controllers "go-crud/controllers/user"
	"go-crud/helpers"

	"github.com/gin-gonic/gin"
)

// InitializeUserRoutes sets up the user routes
func InitializeUserRoutes(r *gin.Engine) {
	routes := []helpers.RouteConfig{
		{Method: "POST", Path: "/login", HandlerFunc: auth.Login},
		{Method: "GET", Path: "/users", HandlerFunc: controllers.GetUsers},
		{Method: "GET", Path: "/users/:id", HandlerFunc: controllers.GetUser},
		{Method: "POST", Path: "/users", HandlerFunc: controllers.CreateUser},
		{Method: "PUT", Path: "/users/:id", HandlerFunc: controllers.UpdateUser},
		{Method: "DELETE", Path: "/users/:id", HandlerFunc: controllers.DeleteUser},
	}

	helpers.RegisterRoutes(r, routes)
}
