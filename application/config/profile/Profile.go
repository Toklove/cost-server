package profile

import (
	"fiber/application/model/Config"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Profile Config.Yaml

func Main() {
	config, _ := ioutil.ReadFile("./config.yaml")
	err := yaml.Unmarshal(config, &Profile)
	if err != nil {
		fmt.Println(err, "配置文件读取失败")
		return
	}
}
