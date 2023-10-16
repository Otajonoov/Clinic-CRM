package v1

import (
	"gitlab.com/clinic-crm/api-gateway/api/models"
	"gitlab.com/clinic-crm/api-gateway/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Error Codes
const (
	ErrorSuccessCode                  = 1000
	ErrorCodeAccessTokenExpired       = 1001
	ErrorCodeRefreshTokenExpired      = 1002
	ErrorCodeNotFound                 = 1003
	ErrorCodeInvalidJson              = 1004
	ErrorCodeWrongPassword            = 1005
	ErrorCodeInternal                 = 1006
	ErrorCodeBadRequest               = 1007
	ErrorCodeUnauthorized             = 1008
	ErrorCodeNotAllowed               = 1009
	ErrorCodeOtpIncorrect             = 1010
	ErrorCodeImageSizeExceed          = 1011
	ErrorCodeImageExtensionNotAllowed = 1012
)

func HandleInternalWithMessage(c *gin.Context, l *logger.Logger, err error, message string, args ...interface{}) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.DefaultResponse{
			ErrorCode:    ErrorCodeInternal,
			ErrorMessage: "Oops something went wrong",
		})
		return true
	}

	return false
}

func HandleDatabaseLevelWithMessage(c *gin.Context, l *logger.Logger, err error, message string, args ...interface{}) bool {
	status_err, _ := status.FromError(err)
	if err != nil {
		errorCode := ErrorCodeInternal
		statuscode := http.StatusInternalServerError
		message := status_err.Message()
		switch status_err.Code() {
		case codes.NotFound:
			errorCode = ErrorCodeNotFound
			statuscode = http.StatusNotFound
		case codes.Unknown:
			errorCode = ErrorCodeInternal
			statuscode = http.StatusBadRequest
			message = "Ooops something went wrong"
		case codes.Aborted:
			errorCode = ErrorCodeBadRequest
			statuscode = http.StatusBadRequest
		case codes.InvalidArgument:
			errorCode = ErrorCodeBadRequest
			statuscode = http.StatusBadRequest
		}

		c.AbortWithStatusJSON(statuscode, models.DefaultResponse{
			ErrorCode:    errorCode,
			ErrorMessage: message,
		})
		return true
	}
	return false
}

func HandleBadRequestErrWithMessage(c *gin.Context, l *logger.Logger, err error, message string, args ...interface{}) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.DefaultResponse{
			ErrorCode:    ErrorCodeBadRequest,
			ErrorMessage: "Please enter right information",
		})
		return true
	}
	return false
}
