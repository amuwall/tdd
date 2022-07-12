package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var globalConfig *Template

func Init(path string) (err error) {
	globalConfig, err = load(path)
	return
}

func Get() *Template {
	return globalConfig
}

func load(path string) (config *Template, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &Template{}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return
	}

	return
}
