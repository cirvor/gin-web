package middleware

import (
	"errors"
	"net/http"
	"web/utils"

	"github.com/gin-gonic/gin"
)

func HandleErrors() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			var err error
			if r := recover(); r != nil {
				switch x := r.(type) {
				case error:
					err = x
				case string:
					err = errors.New(x)
				default:
					err = errors.New("unknown error")
				}
				utils.Error(context, http.StatusInternalServerError, err.Error())
				return
			}
		}()
		context.Next()
	}
}
