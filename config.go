package miniweb

import(
	C "github.com/Unknwon/goconfig"
)

// 用于处理项目设置
type Config struct {
	cf *C.ConfigFile
}

// 初始化项目目录下的配置文件
func InitConfig() *Config {
	config := new(Config)

	// 如果用户没有指定配置文件，那么么配置文件不存在则忽略错误
	cf, err := C.LoadConfigFile("./config.ini")
	if err != nil {
		config.cf = nil
	} else {
		config.cf = cf
	}

	return config

}

// 获取配置
func (c Config) Get(section, key string) string {
	if c.cf == nil {
		return ""
	}
	value, err := c.cf.GetValue(section, key)
	if err != nil {
		panic("Configuration items that do not exist!")
	}
	return value
}
