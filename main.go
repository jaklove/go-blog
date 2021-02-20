package main

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/internal/model"
	"go-blog/internal/router"
	"go-blog/pkg/logger"
	"go-blog/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)



// @title  博客系统
// @version 1.0
// @description Go语言编程之旅:一起用Go做项目
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	newRouter := router.NewRouter()
	S := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        newRouter,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	global.Logger.Infof("%s: go-programming-tour-books/%s","zhourenjie","blog-service")
	S.ListenAndServe()
}

//初始化对应的对象
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setUpSetting: %v", err)
	}

	err = setUpLogger()
	if err != nil{
		log.Fatalf("init.setUpLogger err: %v",err)
	}
	err = setUpDBEngine()
	if err != nil{
		log.Fatalf("init.setUpDBEngine err: %v",err)
	}
}

func setupSetting() error {
	newSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = newSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = newSetting.ReadSection("Database", &global.DataBaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setUpDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DataBaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setUpLogger()error  {
	filename := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: filename,
		MaxSize: 600,
		MaxAge: 10,
		LocalTime: true,
	},"",log.LstdFlags).WithCaller(2)
	return  nil
}