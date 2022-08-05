package main

import (
	"github.com/Gophberg/Store/internal/app/Store"
	"log"
)

func main() {
	if err := Store.Start(); err != nil {
		log.Fatal(err)
	}
}
