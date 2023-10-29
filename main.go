package main

import (
	"api/src/router"
)

func main() {
	r := router.Gerar()
	r.Run()
}
