package main

import (
	"github.com/adnanbrq/slugify/internal/db"
	"github.com/adnanbrq/slugify/internal/http"
)

func main() {
	if err := db.Connect(); err != nil {
		panic(err)
	}

	rest := http.NewRestServer()
	rest.Init()
	if err := rest.Boot(); err != nil {
		panic(err)
	}
}
