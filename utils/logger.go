package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

var (
	Log *logrus.Logger
)
func init() {
	Log = logrus.New()
	// 设置日志级别为xx以及以上
	Log.SetLevel(logrus.DebugLevel)
	//JSON在生产中通常只在使用Splunk或Logstash等工具进行日志聚合时才有用。
	// 设置日志格式为json格式
	// Log.SetFormatter(&logrus.JSONFormatter{
	// 	// PrettyPrint: true,//格式化json
	// 	TimestampFormat: "2006-01-02 15:04:05",//时间格式化
	// })
	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",//时间格式化
	})
	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	Log.SetOutput(os.Stdout)
	//Log.SetReportCaller(true)
	// 初始化一些公共参数
	//loginit:=Log.WithFields(logrus.Fields{
	//	"animal": "walrus",
	//})
	//输出日志
	Log.Info("init logger.......")

}