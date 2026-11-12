package main

import (
	"log"

	"github.com/yourorg/nebula-bridge/internal/config"
	"github.com/yourorg/nebula-bridge/internal/db"
	"github.com/yourorg/nebula-bridge/internal/repository"
)

func main() {
	cfg := config.Load()
	gdb, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	if err := gdb.AutoMigrate(&repository.Node{}, &repository.Event{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}
	log.Println("migration complete")
}
