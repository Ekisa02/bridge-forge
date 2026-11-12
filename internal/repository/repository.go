package repository

import (
	"time"

	"gorm.io/gorm"
)

// Node represents a peer or service in the infrastructure
type Node struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Address   string         `gorm:"size:255;not null;uniqueIndex" json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Event represents an incoming chain event or relay telemetry
type Event struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	NodeID    uint           `gorm:"index;not null" json:"node_id"`
	Type      string         `gorm:"size:100;not null" json:"type"`
	Payload   string         `gorm:"type:text" json:"payload"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// NodeRepository defines node-related persistence operations
type NodeRepository interface {
	Create(node *Node) error
	List() ([]Node, error)
	GetByID(id uint) (*Node, error)
}

// EventRepository defines event-related persistence operations
type EventRepository interface {
	Create(event *Event) error
	List(nodeID *uint) ([]Event, error)
}

// GORM implementations

type gormNodeRepo struct {
	db *gorm.DB
}

func NewGormNodeRepository(db *gorm.DB) NodeRepository {
	return &gormNodeRepo{db: db}
}

func (r *gormNodeRepo) Create(node *Node) error {
	return r.db.Create(node).Error
}

func (r *gormNodeRepo) List() ([]Node, error) {
	var nodes []Node
	if err := r.db.Find(&nodes).Error; err != nil {
		return nil, err
	}
	return nodes, nil
}

func (r *gormNodeRepo) GetByID(id uint) (*Node, error) {
	var node Node
	if err := r.db.First(&node, id).Error; err != nil {
		return nil, err
	}
	return &node, nil
}

// Events

type gormEventRepo struct {
	db *gorm.DB
}

func NewGormEventRepository(db *gorm.DB) EventRepository {
	return &gormEventRepo{db: db}
}

func (r *gormEventRepo) Create(event *Event) error {
	return r.db.Create(event).Error
}

func (r *gormEventRepo) List(nodeID *uint) ([]Event, error) {
	var events []Event
	q := r.db.Order("created_at desc")
	if nodeID != nil {
		q = q.Where("node_id = ?", *nodeID)
	}
	if err := q.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
