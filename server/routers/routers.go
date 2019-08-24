package routers

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

// InitRouter   routerを初期化する
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("api/v1")
	{
		enemies := api.Group("/enemy")
		{
			enemies.GET("/", controllers.ReadEnemies)
			enemies.GET("/:id/kill", controllers.KillEnemy)
		}
	}

	return router
}
