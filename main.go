package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/routers"
	"bluebell/settings"
	"fmt"
)

func main() {
	//第一步,加载配置文件
	if err := settings.Init(); err != nil {
		//加载配置文件出错
		fmt.Printf("load config failed ,err is %v", err)
		return
	}
	//第二步加载日志文件
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		//初始化日志文件出错
		fmt.Printf("init logger failed ,err is %v", err)
		return
	}
	//第三步,mysql初始化
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed ,err is %v", err)
		return
	}
	defer mysql.Close() //程序退出关闭数据库
	//第四步,redis初始化
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("redis init failed ,err is %v", err)
		return
	}
	defer redis.Close()

	//雪花算法生成分布式ID
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed,err:%v\n", err)
		return
	}

	//注册路由 SetupRouter返回值为gin.new()实例
	router := routers.SetupRouter(settings.Conf.Mode)
	err := router.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed,err is %v", err)
		return
	}
}
