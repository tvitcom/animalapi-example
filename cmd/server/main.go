package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tvitcom/animalapi-example/internal/routes"
	"log"
)
const entryPoint = "127.0.0.1:3000"

func main() {
	log.Print("Starting the web app on entry point:" + entryPoint)
	r := gin.Default()
	r.LoadHTMLGlob("./web/templates/**/*")
	app := routes.GetRoutes(r)
	app.Run(entryPoint)
}
