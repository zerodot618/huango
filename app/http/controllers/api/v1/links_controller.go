package v1

import (
	"zerodot618/huango/app/models/link"
	"zerodot618/huango/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	response.Data(c, link.AllCached())
}
