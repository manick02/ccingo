package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Timeout struct {
	Server int `yaml:"server"`
	Read   int `yaml:"read"`
	Write  int `yaml:"write"`
	Idle   int `yaml:"idle"`
}
type ServerConfig struct {
	Timeout Timeout `yaml:"timeout"`
	Host    string  `yaml:"host"`
	Port    int     `yaml:"port"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	//Host int `yaml:"host"`
}

func NewYamlConfig(fileName string) (*Config, error) {
	cin := &Config{}

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("yamlFile read err #%v  at path %v", err, fileName)
	}

	// Read data from yaml file
	err = yaml.Unmarshal(yamlFile, cin)
	if err != nil {
		log.Fatal("Unable to decode config file", err)
	}

	fmt.Println(cin)
	return cin, nil
	//
	//
	//config := &Config{}
	//file,err := ioutil.ReadFile(fileName)
	//if err!=nil {
	//	return nil,err
	//}
	//err = yaml.Unmarshal(file, config)
	//fmt.Println(config)
	//fmt.Println("----")
	//if err != nil {
	//	fmt.Println("Error happened while parsing file",err)
	//	return nil, err
	//}
	//
	//fmt.Println(*config)
	//return config,nil
}

//host: 127.0.0.1
//port: 8080
//timeout:
//server: 30
//read: 15
//write: 10
//idle: 5
