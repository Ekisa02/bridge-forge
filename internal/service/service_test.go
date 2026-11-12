package service_test

import (
	"testing"

	"github.com/yourorg/nebula-bridge/internal/db"
	"github.com/yourorg/nebula-bridge/internal/repository"
	"github.com/yourorg/nebula-bridge/internal/service"
)

func setupService(t *testing.T) *service.Service {
	cfg := &struct{ DatabaseURL string }{DatabaseURL: ""}
	// connect using sqlite fallback (file nebula_test.db)
	_ = cfg
	gdb, err := db.Connect(&struct{
		Port string
		DatabaseURL string
		LogLevel string
	}{DatabaseURL: ""})
	if err != nil {
		t.Fatalf("db connect: %v", err)
	}
	// Use an in-memory sqlite for faster tests
	// NOTE: db.Connect falls back to nebula.db; using same DB file is fine for this unit test context
	_ = gdb
	rnode := repository.NewGormNodeRepository(gdb)
	eventRepo := repository.NewGormEventRepository(gdb)
	_ = gdb.AutoMigrate(&repository.Node{}, &repository.Event{})
	return service.NewService(rnode, eventRepo)
}

func TestCreateListNode(t *testing.T) {
	svc := setupService(t)
	n, err := svc.CreateNode("test-node", "0xabc")
	if err != nil {
		t.Fatalf("create node: %v", err)
	}
	if n.ID == 0 {
		t.Fatalf("expected node id set")
	}
	nodes, err := svc.ListNodes()
	if err != nil {
		t.Fatalf("list nodes: %v", err)
	}
	if len(nodes) == 0 {
		t.Fatalf("expected >=1 nodes")
	}
}

func TestCreateEvent(t *testing.T) {
	svc := setupService(t)
	n, err := svc.CreateNode("ev-node", "0xdef")
	if err != nil {
		t.Fatalf("create node: %v", err)
	}
	e, err := svc.CreateEvent(n.ID, "relay", "{\"payload\":true}")
	if err != nil {
		t.Fatalf("create event: %v", err)
	}
	if e.ID == 0 {
		t.Fatalf("expected event id set")
	}
}
