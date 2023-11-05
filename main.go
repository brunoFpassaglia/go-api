package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
)

func main() {
	config.LoadConfig()
	fmt.Println(config.Port)
	r := router.Gerar()
	r.Run(fmt.Sprintf(":%d", config.Port))
}
