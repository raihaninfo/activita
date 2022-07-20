package main

import (
	"github.com/raihaninfo/activita/models"
	"github.com/raihaninfo/activita/router"
)

const (
	PORT int = 8080
)

func main() {
	db := models.Init()
	router.App(PORT, db)
}
