package tinyconfig

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	VariableOne   string `yaml:"variable_one"`
	VariableTwo   string `yaml:"variable_two"`
}

func readFile(path string, cfg *config) {
	cffilepath := path + "/config.yml"
	f, err := os.Open(cffilepath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadConfig(path string) config {
	cfg := config{}
	readFile(path, &cfg)
	return cfg
}
