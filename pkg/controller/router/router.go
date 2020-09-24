package router

import (
	"github.com/gin-gonic/gin"
	"github.com/otera05/go-webserver-sample/pkg/controller/handler"
)

// NewRouter ...
func NewRouter(r *gin.Engine, handlers *handler.Handlers) {
	r.GET("/", handlers.Hello.Get)
}
