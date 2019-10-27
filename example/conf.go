/**
	goseed 的 configuration 功能允许你通过两种方式来将配置文件转换为对应的结构体：
	1. ini
	2. yaml
 */
package example

import (
	"fmt"
	goseed "powhole.com/go-seed"
)

type Server struct {
	Host string
	Port int
}

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
}

func main() {
	configureByIni()
	configureByYaml()
}

// 通过 ini 文件进行配置
func configureByIni() {
	server := &Server{}

	m := goseed.ConfigurationMap{}
	m["server"] = server
	goseed.ConfigureByIni("conf/config.ini", m)

	fmt.Println(server)
}

// 通过 yml 文件进行配置
func configureByYaml() {
	config := &Config{}
	goseed.ConfigureByYaml("conf/config.yml", config)

	fmt.Println(config)
}