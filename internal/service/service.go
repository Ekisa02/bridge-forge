package service

import (
	"errors"

	"github.com/yourorg/nebula-bridge/internal/repository"
)

type Service struct {
	nodeRepo  repository.NodeRepository
	eventRepo repository.EventRepository
}

func NewService(n repository.NodeRepository, e repository.EventRepository) *Service {
	return &Service{nodeRepo: n, eventRepo: e}
}

// CreateNode validates and creates a node
func (s *Service) CreateNode(name, address string) (*repository.Node, error) {
	if name == "" || address == "" {
		return nil, errors.New("name and address are required")
	}
	node := &repository.Node{Name: name, Address: address}
	if err := s.nodeRepo.Create(node); err != nil {
		return nil, err
	}
	return node, nil
}

func (s *Service) ListNodes() ([]repository.Node, error) {
	return s.nodeRepo.List()
}

// CreateEvent attaches an event to a node
func (s *Service) CreateEvent(nodeID uint, typ, payload string) (*repository.Event, error) {
	if typ == "" {
		return nil, errors.New("type is required")
	}
	// verify node exists
	if _, err := s.nodeRepo.GetByID(nodeID); err != nil {
		return nil, err
	}
	e := &repository.Event{NodeID: nodeID, Type: typ, Payload: payload}
	if err := s.eventRepo.Create(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (s *Service) ListEvents(nodeID *uint) ([]repository.Event, error) {
	return s.eventRepo.List(nodeID)
}
