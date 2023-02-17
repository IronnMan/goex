package auth

import (
	"github.com/gin-gonic/gin"
	v1 "goex/app/http/controllers/api/v1"
	"goex/pkg/captcha"
	"goex/pkg/logger"
	"goex/pkg/response"
)

type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha show captcha image
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {

	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()

	logger.LogIf(err)

	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
