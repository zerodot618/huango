package v1

import (
	"strconv"
	"time"
	"zerodot618/huango/app/requests"
	"zerodot618/huango/pkg/auth"
	"zerodot618/huango/pkg/cache"
	"zerodot618/huango/pkg/config"
	"zerodot618/huango/pkg/response"
	"zerodot618/huango/pkg/shortener"

	"github.com/gin-gonic/gin"
)

type ShortenerController struct {
	BaseAPIController
}

func (ctrl *ShortenerController) CreateShortURL(c *gin.Context) {
	request := requests.ShortenerRequest{}
	if ok := requests.Validate(c, &request, requests.ShortenerCreate); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	if currentUser.ID != request.UserId {
		response.Unauthorized(c, "用户 id 不一致")
	}

	shortUrl := shortener.GenerateShortURL(request.LongURL, strconv.FormatUint(request.UserId, 10))
	cache.Set(shortUrl, request.LongURL, 6*time.Hour)

	response.JSON(c, gin.H{
		"short_url": config.GetString("app.url") + shortUrl,
	})
}

func (ctrl *ShortenerController) ReturnLongURL(c *gin.Context) {
	shortUrl := c.Param("short-url")
	initialUrl := cache.Get(shortUrl)
	response.JSON(c, gin.H{
		"short_url": config.GetString("app.url") + shortUrl,
		"long_url":  initialUrl,
	})
}
