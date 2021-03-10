package middleware

import (
	"github.com/gin-gonic/gin"
	"go-blog/pkg/app"
	"go-blog/pkg/errcode"
	"go-blog/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface)gin.HandlerFunc  {
	return func(context *gin.Context) {
		key := l.Key(context)
		if bucket,ok := l.GetBucket(key);ok{
			count := bucket.TakeAvailable(1)
			if count == 0{
				reponse := app.NewResponse(context)
				reponse.ToErrorResponse(errcode.TooManyRequests)
				return
			}
		}
		context.Next()
	}
}
