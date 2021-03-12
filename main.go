package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/internal/model"
	"go-blog/internal/router"
	"go-blog/pkg/logger"
	"go-blog/pkg/setting"
	"go-blog/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
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

	//平滑重启
	go func() {
		if err := S.ListenAndServe();err != nil && err != http.ErrServerClosed{
			log.Fatalf("s.ListenAndServe err:%v",err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	log.Println("shuting down server...")

	//最大时间控制，通知该服务端它有5s的时间来处理原有的请求
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)

	defer cancel()
	if err := S.Shutdown(ctx);err != nil{
		log.Fatal("Server forced to shutdown:",err)
	}

	log.Println("Server exiting")


	//global.Logger.Infof("%s: go-programming-tour-books/%s","zhourenjie","blog-service")
	//S.ListenAndServe()
}

//初始化对应的对象
func init() {
	setupFlag()
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

	err = setupTracer()
	if err != nil{
		log.Fatalf("init.setupTracer err: %v",err)
	}
}

var (
	port string
	runMode string
	config string
)

func setupFlag()error  {
	flag.StringVar(&port,"port","","启动端口")
	flag.StringVar(&runMode,"mode","","启动模式")
	flag.StringVar(&config,"config","configs/","指定要使用的配置文件路径")
	flag.Parse()

	fmt.Println("port：   ",port)
	return nil
}



func setupTracer()error  {
	jaegerTracer,_,err := tracer.NewJaegerTracer("blog-service","192.168.56.100:6831")
	if err != nil{
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

func setupSetting() error {
	newSetting, err := setting.NewSetting(strings.Split(config,",")...)
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

	err = newSetting.ReadSection("MusicDataBase", &global.MusicSetting)
	if err != nil {
		return err
	}


	err = newSetting.ReadSection("JWT",&global.JWTSetting)
	if err != nil{
		return err
	}

	err = newSetting.ReadSection("Email",&global.EmailSetting)
	if err != nil{
		return err
	}

	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	fmt.Println("port:",port)
	if port != ""{
		global.ServerSetting.HttpPort = port
	}
	if runMode != ""{
		global.ServerSetting.RunMode = runMode
	}

	return nil
}

func setUpDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DataBaseSetting)
	if err != nil {
		return err
	}

	global.MusicDBEngine,err = model.NewMusicDBEngine(global.MusicSetting)
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