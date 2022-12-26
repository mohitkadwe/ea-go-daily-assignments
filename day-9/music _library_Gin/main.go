package songs

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start with Go Gin Gonic Framework")

	route := gin.New()

	route.Use(gin.Logger())

	route.Use(gin.Recovery())

	route.GET("/songs?id=id", getSong)

	route.POST("/songs", createSong)

	route.PUT("/songs", updateSong)

	http.ListenAndServe("localhost:8080", route)
}
