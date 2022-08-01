package main

import (
	"github.com/raihaninfo/activita/models"
	"github.com/raihaninfo/activita/router"
)

const (
	PORT int = 8080
)

func main() {
	models.Init()
	router.App(PORT)
}
