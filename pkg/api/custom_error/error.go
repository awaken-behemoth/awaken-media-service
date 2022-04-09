package custom_error

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ApiErrorResponse struct {
	error ApiError `json:"error"`
}

func NewApiError(code int, message string) ApiError {
	return ApiError{
		code:      code,
		message:   message,
		status:    http.StatusText(code),
		timestamp: time.Now().Format(time.RFC850),
	}
}

type ApiError struct {
	code      int    `json:"code"`
	status    string `json:"status"`
	message   string `json:"message"`
	timestamp string `json:"timestamp"`
}

func (apiError ApiError) Error() string {
	return apiError.message
}

func NewHandler() gin.HandlerFunc {

	return handleError
}

func handleError(context *gin.Context) {

	context.Next()

	detectedErrors := context.Errors.ByType(gin.ErrorTypeAny)

	if len(detectedErrors) > 0 {
		err := detectedErrors[0].Err
		var parsedError *ApiError

		switch err.(type) {

		case *ApiError:
			print("I matched the api error")
			parsedError = err.(*ApiError)
			break

		default:
			print("I matched the dfault")
			temp := NewApiError(http.StatusInternalServerError, err.Error())
			parsedError = &temp
		}

		fmt.Printf("t1: %T\n", err)

		//print(parsedError.code)
		context.JSON(parsedError.code, parsedError.Error())
		//context.AbortWithStatusJSON(parsedError.code, gin.H{"error": parsedError.Error()})
		return
	}

}
