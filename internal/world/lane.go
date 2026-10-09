package world

import (
	"math/rand"

	"example.ezboot/myna/config"
)

// Lane 是一条寂静的小路
type Lane struct {
	cfg *config.World
}

// Via 路过一条寂静的小路
func (l *Lane) Via(text string) string {
	if rand.Int31n(2) == 0 {
		return l.cfg.Mysterious + " " + text
	}
	return text
}

// NewLane 创建一条寂静的小路
func NewLane(cfg *config.World) *Lane {
	return &Lane{cfg: cfg}
}
