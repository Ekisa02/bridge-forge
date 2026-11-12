package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yourorg/nebula-bridge/internal/config"
	"github.com/yourorg/nebula-bridge/internal/db"
	"github.com/yourorg/nebula-bridge/internal/handlers"
	"github.com/yourorg/nebula-bridge/internal/repository"
	"github.com/yourorg/nebula-bridge/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	gdb, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("failed connecting to DB: %v", err)
	}

	// Auto-migrate schema (safe for initial development)
	if err := gdb.AutoMigrate(&repository.Node{}, &repository.Event{}); err != nil {
		log.Fatalf("auto-migrate failed: %v", err)
	}

	nodeRepo := repository.NewGormNodeRepository(gdb)
	eventRepo := repository.NewGormEventRepository(gdb)

	svc := service.NewService(nodeRepo, eventRepo)

	r := gin.Default()
	h := handlers.NewHandler(svc)

	r.GET("/health", handlers.Health)
	r.POST("/nodes", h.CreateNode)
	r.GET("/nodes", h.ListNodes)
	r.POST("/events", h.CreateEvent)
	r.GET("/events", h.ListEvents)

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":" + port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}

	os.Exit(0)
}
