package models

import (
	"time"

	"github.com/google/uuid"
)

type TableTopGame struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	GameDetail GameDetail `json:"gameDetail"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

type GameDetail struct {
	Description string `json:"description"`
	Rules       string `json:"rules"`
	Points      int    `json:"points"`
}

type TableTopGameReq struct {
	Name       string     `json:"name"`
	GameDetail GameDetail `json:"gameDetail"`
}

type TableTopGameResp struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	GameDetail GameDetail `json:"gameDetail"`
}
