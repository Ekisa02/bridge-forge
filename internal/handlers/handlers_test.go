package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yourorg/nebula-bridge/internal/db"
	"github.com/yourorg/nebula-bridge/internal/handlers"
	"github.com/yourorg/nebula-bridge/internal/repository"
	"github.com/yourorg/nebula-bridge/internal/service"
)

func setupRouter(t *testing.T) *gin.Engine {
	gdb, err := db.Connect(&struct{
		Port string
		DatabaseURL string
		LogLevel string
	}{})
	if err != nil {
		t.Fatalf("db connect: %v", err)
	}
	_ = gdb.AutoMigrate(&repository.Node{}, &repository.Event{})
	nodeRepo := repository.NewGormNodeRepository(gdb)
	eventRepo := repository.NewGormEventRepository(gdb)
	svc := service.NewService(nodeRepo, eventRepo)
	h := handlers.NewHandler(svc)
	g := gin.New()
	g.POST("/nodes", h.CreateNode)
	g.GET("/nodes", h.ListNodes)
	g.POST("/events", h.CreateEvent)
	return g
}

func TestCreateNodeHandler(t *testing.T) {
	g := setupRouter(t)
	w := httptest.NewRecorder()
	body := map[string]string{"name": "h-node", "address": "0xhandler"}
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/nodes", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	g.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 got %d", w.Code)
	}
}
