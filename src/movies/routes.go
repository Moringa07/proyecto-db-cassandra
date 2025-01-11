package movies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registra las rutas para el grupo de películas
func RegisterRoutes(router *gin.Engine) {
	group := router.Group("/movies")

	group.GET("/", func(c *gin.Context) {
		movies, err := getAllMoviesController()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch movies"})
			return
		}
		c.JSON(http.StatusOK, movies)
	})
}

