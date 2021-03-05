package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-blog/pkg/app"
	"go-blog/pkg/errcode"
)

func JWT()gin.HandlerFunc  {
	return func(context *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)

		if s, exists := context.GetQuery("token");exists{
			token = s
		}else {
			token = context.GetHeader("token")
		}

		if token == ""{
			ecode = errcode.InvalidParams
		}else {
			_, err := app.ParseToken(token)
			if err != nil{
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success{
			response := app.NewResponse(context)
			response.ToErrorResponse(ecode)
			context.Abort()
			return
		}
		context.Next()
	}

}
