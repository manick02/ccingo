package main

import "fmt"

func main() {
	filename:="/Users/8420/workspace/goprojects/src/github.com/manick02/ccingo/yamlp/config.yaml"
	fmt.Println(filename)
	config,err:=NewYamlConfig(filename)
    if (err!=nil) {
	 fmt.Println(config)
    }
    fmt.Println(err)
}
