package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamikazechaser/disposable-email-cache/internal/tasks"
	"github.com/patrickmn/go-cache"
)

var (
	loadedCache *cache.Cache
)

func Start() error {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	router := gin.Default()

	router.GET("/update-cache", updateCacheHandler)
	router.GET("/check/:email", checkDisposableHandler)

	loadedCache = tasks.LoadCache()
	return router.Run(":5000")
}

func updateCacheHandler(c *gin.Context) {
	tasks.DownloadData()
	loadedCache.Flush()
	loadedCache = tasks.LoadCache()

	c.JSON(http.StatusOK, gin.H{
		"cacheUpdated": true,
	})
}

func checkDisposableHandler(c *gin.Context) {
	email := c.Param("email")

	_, found := loadedCache.Get(email)

	c.JSON(http.StatusOK, gin.H{
		"disposable": found,
	})
}
