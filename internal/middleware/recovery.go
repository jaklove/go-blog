package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/pkg/Email"
	"go-blog/pkg/app"
	"go-blog/pkg/errcode"
	"time"
)

func Recovery()gin.HandlerFunc  {
	defaultMailer := Email.NewEmail(&Email.SMTPInfo{
		Host: global.EmailSetting.Host,
		Port: global.EmailSetting.Port,
		IsSSL: global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		From: global.EmailSetting.From,
	})

	return func(context *gin.Context) {
		defer func() {
			if err := recover();err != nil{
				global.Logger.WithCallersFrames().Errorf("panic recover err: %v",err)
				err := defaultMailer.SendMail(global.EmailSetting.To,fmt.Sprintf("异常抛出，发生时间：%d",time.Now().Unix()),fmt.Sprintf("错误信息: %v",err))
				if err != nil{
					global.Logger.Errorf("mail.sendmail:%v",err)
				}
				app.NewResponse(context).ToErrorResponse(errcode.ServerError)
				context.Abort()
			}
		}()
		context.Next()
	}
}
