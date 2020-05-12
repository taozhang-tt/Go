package app

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLog() {
	G.Log = logrus.New()
	G.Log.SetFormatter(&logrus.JSONFormatter{})	//设置输出格式
	G.Log.SetOutput(os.Stdout)	//设置输出路径
	//G.Log.SetLevel(logrus.WarnLevel)	//设置日志等级
	//G.Log.SetReportCaller(true)	//add the calling method as a field
}
