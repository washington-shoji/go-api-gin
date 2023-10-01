package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type ExpHandler struct {
	Database *sql.DB
}

type DynamicStruct struct {
	ID        uuid.UUID   `json:"id"`
	Data      interface{} `json:"data"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt *time.Time  `json:"updatedAt"`
	DeletedAt *time.Time  `json:"deletedAt"`
}

func NewExpHandler(Db *sql.DB) *ExpHandler {
	return &ExpHandler{
		Database: Db,
	}
}

func (handler *ExpHandler) Exp(ctx *gin.Context) {
	raw, err := ctx.GetRawData()
	if err != nil {
		fmt.Println("err 1", err)
		return
	}
	var buf bytes.Buffer
	if err := json.Compact(&buf, raw); err != nil {
		fmt.Println("err 2", err)
		return
	}
	data := buf.Bytes()
	//fmt.Println("compact json", string(data))

	//var target map[string]any
	dynamicStruct := DynamicStruct{}

	if err := json.Unmarshal([]byte(data), &dynamicStruct.Data); err != nil {
		log.Fatalf("Unable to marshal JSON due to %s", err)
	}

	id := uuid.New()
	time := time.Now().UTC()

	db := handler.Database

	query := `INSERT INTO table_exp (id, exp_json, created_at)
	VALUES ($1, $2, $3)
	RETURNING *
	`

	result, err := db.Query(query, id, data, time)
	if err != nil {
		log.Fatalf("Unable to save JSON data due to %s", err)
	}

	fmt.Println("result >>> ", result)

	//fmt.Println("data >>> ", data)
	//fmt.Println("target >>> ", target)

	dynamicIterate := dynamicStruct.Data.(map[string]interface{})
	for k, v := range dynamicIterate {
		fmt.Printf("k: %s, v: %v\n", k, v)
	}

	// for k, v := range target {
	// 	fmt.Printf("k: %s, v: %v\n", k, v)
	// }
}

// func (handler *ExpHandler) Exp(ctx *gin.Context) {
// 	raw, err := ctx.GetRawData()
// 	if err != nil {
// 		fmt.Println("err 1", err)
// 		return
// 	}
// 	var data Data
// 	buf := []byte(raw)

// 	if err := json.Unmarshal(buf, &data); err != nil {
// 		fmt.Println("err 2", err)
// 	}

// 	fmt.Println("data", data)
// }
