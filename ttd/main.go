package main

import (
	"log"

	"github.com/zackartz/ttd/api"
)

func main() {
	server := api.Server{}

	log.Println("[TTD] Starting...")
	server.Initialize()
}
