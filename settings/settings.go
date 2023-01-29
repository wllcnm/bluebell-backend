package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper" //viper是用来加载配置文件的框架，有了它就不用自己手写反射了。
)

var Conf = new(AppConfig)

type AppConfig struct {
	Mode         string `mapstructure:"mode"` //tag表示的是与配置文件中的属性一一对应
	Port         int    `mapstructure:"port"`
	Name         string `mapstructure:"name"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int    `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func Init() error {
	//加载配置文件,以main.go为起始文件
	viper.SetConfigFile("./conf/config.yaml")

	/*
		WatchConfig()和OnConfigChange()是一套函数，一个用来监控，一个用来设置预案
	*/
	//对配置文件进行监视
	viper.WatchConfig()

	//如果配置文件被修改
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件被修改了")
		_ = viper.Unmarshal(&Conf) //er.Unmarshal(&u)反序列化，指的是将配置文件的内容加载到结构体中，记得用指针
	})
	err := viper.ReadInConfig()
	if err != nil {
		//fmt.Printf("readconfig failed,err is%v\n",err)
		panic(fmt.Errorf("readconfig failed,err is%v\n", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal failed ,err is %v", err))
	}
	return err
}
