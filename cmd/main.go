package main

import (
	"github.com/bootcamp-go/desafio-cierre-testing/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	if err := r.Run(":18085"); err != nil {
		panic(err)
	}
}
