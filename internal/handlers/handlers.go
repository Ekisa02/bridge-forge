package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourorg/nebula-bridge/internal/service"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

type createNodeReq struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

func (h *Handler) CreateNode(c *gin.Context) {
	var req createNodeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	node, err := h.svc.CreateNode(req.Name, req.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, node)
}

func (h *Handler) ListNodes(c *gin.Context) {
	nodes, err := h.svc.ListNodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

type createEventReq struct {
	NodeID  uint   `json:"node_id" binding:"required"`
	Type    string `json:"type" binding:"required"`
	Payload string `json:"payload"`
}

func (h *Handler) CreateEvent(c *gin.Context) {
	var req createEventReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	e, err := h.svc.CreateEvent(req.NodeID, req.Type, req.Payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, e)
}

func (h *Handler) ListEvents(c *gin.Context) {
	var nodeID *uint
	if v := c.Query("node_id"); v != "" {
		i, err := strconv.Atoi(v)
		if err == nil {
			i64 := uint(i)
			nodeID = &i64
		}
	}
	events, err := h.svc.ListEvents(nodeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}
