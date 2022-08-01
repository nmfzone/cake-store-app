package response

import (
	"github.com/gin-gonic/gin"
	"github.com/nmfzone/privy-cake-store/internal/errors"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func CreatedResponse(ctx *gin.Context, msg string, m interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": msg,
		"data":    m,
	})
}

func OkResponseData(ctx *gin.Context, m interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": m,
	})
}

func OkResponseWith(ctx *gin.Context, msg string, m interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
		"data":    m,
	})
}

func OkResponse(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

func BadRequestResponse(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": msg,
	})

	defer ctx.AbortWithStatus(http.StatusBadRequest)
}

func NewResponseError(ctx *gin.Context, err error) {
	errorCode := http.StatusInternalServerError

	switch err {
	case errors.ErrInternalServerError:
		errorCode = http.StatusInternalServerError
	case errors.ErrNotFound:
		errorCode = http.StatusNotFound
	case errors.ErrConflict:
		errorCode = http.StatusConflict
	}

	ctx.JSON(errorCode, gin.H{
		"message": err.Error(),
	})

	defer ctx.AbortWithStatus(errorCode)
}

func ValidationErrorResponse(ctx *gin.Context, v errors.ValidationError) {
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": "The given data was invalid.",
		"errors":  v.Errors,
	})

	defer ctx.AbortWithStatus(http.StatusUnprocessableEntity)
}
