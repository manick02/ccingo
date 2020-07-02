package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Timeout struct {
	Server int `yaml:"server"`
	Read int `yaml:"read"`
	Write int `yaml:"write"`
	Idle int `yaml:"idle"`
}
type ServerConfig struct {
	Timeout Timeout `yaml:"timeout"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Config struct {
	server ServerConfig `yaml:"server"`
}

func NewYamlConfig(fileName string) (*Config,error) {
	config := &Config{}
	file,err := os.Open(fileName)
	if err!=nil {
		return nil,err
	}
	defer file.Close()
	var d = yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	fmt.Println(*config)

	return config,nil
}

//host: 127.0.0.1
//port: 8080
//timeout:
//server: 30
//read: 15
//write: 10
//idle: 5