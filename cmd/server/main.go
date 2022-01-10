package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tvitcom/animalapi-example/internal/routes"
	"log"
)
const entryPoint = "127.0.0.1:3000"

func main() {
	log.Print("Starting the webapp on:" + entryPoint)
	
	r := gin.New()
	r.LoadHTMLGlob("./web/templates/**/*")

	// setup limits:
	r.MaxMultipartMemory = 8 << 20  // 8 MiB
	
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	
	app := routes.GetRoutes(r)
	app.Run(entryPoint)
}
