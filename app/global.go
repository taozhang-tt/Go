package app

import (
	"github.com/sirupsen/logrus"
)

var G *Global

// G 表示全局变量
type Global struct {
	Log *logrus.Logger
}

func InitGlobalVar() {
	// 初始化全局结构体
	G = new(Global)

	// 初始化配置

	// 初始化环境

	// 初始化日志
	InitLog()
}
