package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"producer/routes"
)

func main() {
	fmt.Println("Producer: I'm sentient.")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to the producer.")
	})

	r.POST(routes.POST_USER_ROUTE, routes.PostUser)

	r.Run(":80") // listen and serve on 0.0.0.0:8080

}