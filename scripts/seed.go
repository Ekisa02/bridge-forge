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
	_ = gdb.AutoMigrate(&repository.Node{}, &repository.Event{})

	r := repository.NewGormNodeRepository(gdb)
	// seed sample node
	if _, err := r.List(); err == nil {
		r.Create(&repository.Node{Name: "seed-node", Address: "0xseed"})
	}
	log.Println("seed done")
}
