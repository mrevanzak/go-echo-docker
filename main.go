package main

import (
	"Praktikum/config"
	"Praktikum/route"
)

func main() {
	config.InitDatabase()

	e := route.New()
	e.Start(":8080")
}
