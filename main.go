package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"github.com/Kelniit/Marta/router"
)

func main(){
	// Gin Default
	rute := gin.Default()

	router.ProductRouter(rute)

	rute.Run()
}