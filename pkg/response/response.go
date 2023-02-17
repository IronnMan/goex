package response

import (
	"github.com/gin-gonic/gin"
	"goex/pkg/logger"
	"gorm.io/gorm"
	"net/http"
)

// JSON respond with 200 status code and JSON data
func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// Success respond with 200 status code and preset "operation successful!" JSON data
// called after "change" operation "no specific return data" is executed successfully, such as deleting, changing
// password, changing mobile phone number
func Success(c *gin.Context) {
	JSON(c, gin.H{
		"success": true,
		"message": "successful operation!",
	})
}

// Data respond with 200 status code and JSON data with data key
func Data(c *gin.Context, data interface{}) {
	JSON(c, gin.H{
		"success": true,
		"data":    data,
	})
}

// Created respond with 201 status code and JSON data with data key
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    data,
	})
}

// CreatedJSON respond with 201 status code and JSON data
func CreatedJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

// Abort404 respond with 404 status code, use the default message when no parameter msg is passed
func Abort404(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": defaultMessage("The data does not exist, please make sure the request is correct", msg...),
	})
}

// Abort403 respond with 403 status code, use the default message when no parameter msg is passed
func Abort403(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": defaultMessage("Insufficient permissions, please make sure you have the corresponding permissions", msg...),
	})
}

// Abort500 respond with 500 status code, use the default message when no parameter msg is passed
func Abort500(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": defaultMessage("Internal server, please try again later", msg...),
	})
}

// BadRequest respond with 400 status code, pass the parameter err object, use the default message when no parameter msg is passed
func BadRequest(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": defaultMessage("Request parsing error, please confirm whether the request format is correct."+
			"Please use the multipart header when uploading files, and use JSON format for parameters", msg...),
		"error": err.Error(),
	})
}

// Error respond with 404 or 402 status code, use the default message when no parameter msg is passed
func Error(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)

	if err == gorm.ErrRecordNotFound {
		Abort404(c)
		return
	}

	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": defaultMessage("Request processing failed, please check the value of error", msg...),
		"error":   err.Error(),
	})
}

// ValidationError handle the error that the form validation fails, the returned JSON example:
//
//	{
//	    "errors": {
//	        "email": [
//	            "The email is required, and the parameter name is email",
//	        ]
//	    },
//	    "message": "Request verification failed, please see errors for details"
func ValidationError(c *gin.Context, errors map[string][]string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": "Request verification failed, please see errors for details",
		"errors":  errors,
	})
}

// Unauthorized respond with 401 status code, use the default message when no parameter msg is passed
func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": defaultMessage("Request parsing error, please confirm whether the request format is correct."+
			"Please use the multipart header when uploading files, and use JSON format for parameters", msg...),
	})
}

// defaultMessage built-in helper function to support default parameter default values
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}
	return
}
