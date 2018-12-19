package main

import (
	"fmt"
	"flag"
)

func main() {
	config := flag.String("config", "null", "-config 配置数据")
	flag.Parse()
	fmt.Println("exe_demo.go", "config", *config)
}
