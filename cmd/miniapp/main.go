package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/otera05/go-webserver-sample/pkg/controller/handler"
	"github.com/otera05/go-webserver-sample/pkg/controller/router"
)

func main() {
	handlers := handler.NewHandlers()
	r := gin.Default()
	router.NewRouter(r, handlers)
	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
