package main

import (
	"fmt"
	"project/internal/config"
)

func main() {
	//TODO init config
	cfg := config.MustReadConfigFile()
	fmt.Println(cfg)

	//TODO init logger

	//TODO init storage

	//TODO init router

	//TODO run server

}
