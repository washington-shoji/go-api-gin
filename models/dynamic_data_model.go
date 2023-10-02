package models

import (
	"time"

	"github.com/google/uuid"
)

type DynamicData struct {
	ID        uuid.UUID   `json:"id"`
	Data      interface{} `json:"data"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt *time.Time  `json:"updatedAt"`
	DeletedAt *time.Time  `json:"deletedAt"`
}

type DynamicDataReq struct {
	Data interface{} `json:"data"`
}

type DynamicDataRes struct {
	ID   uuid.UUID   `json:"id"`
	Data interface{} `json:"data"`
}
