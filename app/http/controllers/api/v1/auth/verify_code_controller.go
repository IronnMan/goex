package auth

import (
	"github.com/gin-gonic/gin"
	v1 "goex/app/http/controllers/api/v1"
	"goex/app/requests"
	"goex/pkg/captcha"
	"goex/pkg/logger"
	"goex/pkg/response"
	"goex/pkg/verifycode"
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

// SendUsingEmail send Email verification code
func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {
	request := requests.VerifyCodeEmailRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok {
		return
	}

	err := verifycode.NewVerifyCode().SendEmail(request.Email)

	if err != nil {
		response.Abort500(c, "Failed to send email verification code")
	} else {
		response.Success(c)
	}
}
