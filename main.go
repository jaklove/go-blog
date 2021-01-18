package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/internal/model"
	"go-blog/internal/routers"
	"go-blog/pkg/logger"
	"go-blog/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init()  {
	err := setupSettings()
	if err != nil{
		log.Fatalf("init.setupsetting err: %v",err)
	}
	err = setupLogger()
	if err != nil{
		log.Fatal("init.setupLogger err: %v",err)
	}
}


func main()  {
	global.Logger.Info("%s:go-blog/%s","zhourenjie")
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	fmt.Println("port:",global.ServerSetting.HttpPort)
	s := &http.Server{
		Addr:":"+global.ServerSetting.HttpPort,
		Handler:router,
		ReadTimeout:global.ServerSetting.ReadTimeout,
		WriteTimeout:global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}


func setupSettings()error  {
	setting, err := setting.NewSetTing()
	if err != nil{
		return err
	}

	section := setting.ReadSection("Server", &global.ServerSetting)
	err = section
	if err != nil{
		fmt.Println(global.ServerSetting)
		return  err
	}

	err = setting.ReadSection("App",&global.AppSetting)
	if err != nil{
		return  err
	}

	err = setting.ReadSection("Database",&global.DatabaseSetting)
	if err != nil{
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine()error  {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil{
		return  err
	}
	return  nil
}
func setupLogger()error  {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:global.AppSetting.LogSavePath+"/"+global.AppSetting.LogFileName+global.AppSetting.LogFileExt,
		MaxSize:600,
		MaxAge:10,
		LocalTime:true,
	},"",log.LstdFlags).WitchCaller(2)
	return nil
}