package router

import (
	"fmt"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"koala/src/apis/indexApis"
	"koala/src/ginMiddleWare/logger"
	"koala/src/log"
	"koala/src/util/jsonUtil"
)

func InitRouter(port string) error {
	//修改模式
	gin.SetMode(gin.ReleaseMode)
	//r := gin.Default()
	r := gin.New()
	err := logger.InitRouteLogger()
	if err != nil {
		fmt.Println("Failed to init route logger.")
		return err
	}
	r.Use(logger.GinLogger(logger.Logger))
	//r.Use(logger.GinLogger(logger.Logger, time.RFC3339, true))
	r.Use(logger.GinRecovery(logger.Logger, true))
	r.Use(cors.Default())

	//r.Use(gin.Recovery())
	//r.Use(logger.Logger())
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	//gin.DisableConsoleColor()
	// 记录到文件
	//file := configUtil.GetLogFile()

	//file := "./logs/test.log"
	//os.Create(file)
	//f, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//gin.DefaultWriter = io.MultiWriter(f)
	//logLevel := configUtil.GetLogLevel()

	//gin.DefaultWriter = io.MultiWriter(os.Stdout, &hook)
	//gin.DefaultWriter = io.MultiWriter(os.Stdout, &hook)

	//设定请求url不存在的返回值
	r.NoRoute(jsonUtil.NoResponse)

	r.GET("/", indexApis.DefaultIndexApi)

	//user := r.Group("user")
	//{
	//	user.POST("/login", apis.LoginApi)
	//	user.POST("/updatePassword", apis.UpdatePsdApi)
	//}
	//
	//academy := r.Group("academy")
	//{
	//	academy.POST("/", apis.AddOrUpdateAcademyApi)
	//	academy.GET("/:academyId", apis.FindOneAcademyApi)
	//	academy.DELETE("/:academyId", apis.DeleteAcademyApi)
	//	//分页查询
	//	academy.POST("/search/:page/:limit", apis.SearchAcademyApi)
	//	//academy.POST("/test", apis.TestApi)
	//}
	//
	//attendance := r.Group("attendance")
	//{
	//	attendance.POST("/updateAttendance", apis.UpdateClockState)
	//	attendance.GET("/queryAttendanceCountAndRateOfSection/:sectionId", apis.SearchAttendanceCountAndRateOfSection)
	//	attendance.POST("/queryAttendancesOfStudent/:page/:limit", apis.SearchAttendancesOfStudent)
	//}

	err = r.Run(":" + port)
	if err != nil {
		log.Error("Failed to run go-gin web router, try again please. \n Error message:", zap.Error(err))
		return err
	}

	log.Info("Run go-gin web router successfully.")
	return nil
}
