package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ShortenerRequest struct {
	LongURL string `valid:"long_url" json:"long_url,omitempty"`
	UserId  uint64 `valid:"user_id" json:"user_id,omitempty"`
}

func ShortenerCreate(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"long_url": []string{"required"},
		"user_id":  []string{"required"},
	}
	messages := govalidator.MapData{
		"long_url": []string{
			"required:原始地址必填项",
		},
		"user_id": []string{
			"required:当前用户不存在",
		},
	}
	return validate(data, rules, messages)
}
