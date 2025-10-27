package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zxc7563598/avatar-service/internal/cache"
	"github.com/zxc7563598/avatar-service/internal/generator"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "服务器端口")
	flag.StringVar(&port, "p", "8080", "服务器端口（简写）")
	flag.Parse()
	r := gin.Default()
	// 缓存实例
	memCache := cache.NewMemoryCache()
	r.GET("/avatar", func(c *gin.Context) {
		seed := c.Query("seed")
		if seed == "" {
			c.String(http.StatusBadRequest, "missing seed")
			return
		}
		size := 512
		style := c.DefaultQuery("style", "identicon")
		// 缓存 key
		cacheKey := style + ":" + seed + ":" + strconv.Itoa(size)
		if imgBytes, ok := memCache.Get(cacheKey); ok {
			c.Data(http.StatusOK, "image/png", imgBytes)
			return
		}
		// 生成头像
		var img []byte
		var err error
		switch style {
		case "identicon":
			img, err = generator.GenerateIdenticon(seed, size)
		case "pixel":
			img, err = generator.GeneratePixelAvatar(seed, size)
		default:
			img, err = generator.GenerateIdenticon(seed, size)
		}
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		// 缓存
		memCache.Set(cacheKey, img)
		c.Data(http.StatusOK, "image/png", img)
	})
	r.Run(":" + port)
}
