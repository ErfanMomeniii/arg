package main

import (
	"github.com/erfanmomeniii/arg/internal/app"
	"log"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
