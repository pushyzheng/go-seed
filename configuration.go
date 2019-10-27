package goseed

import (
	"fmt"
	"github.com/go-ini/ini"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigurationMap map[string]interface{}

var cfg *ini.File

func ConfigureByIni(path string, conf ConfigurationMap) {
	var err error
	cfg, err = ini.Load(path)
	if err != nil {
		parseError(path, err)
	}
	for name, p := range conf {
		mapTo(name, p)
	}
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		panic(fmt.Sprintf("Fail to parse struct %v", err))
	}
}

func ConfigureByYaml(path string, conf interface{}) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		parseError(path, err)
	}
}

func parseError(path string, err error) {
	panic(fmt.Sprintf("Fail to parse '%s': %v", path, err))
}
