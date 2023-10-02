package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/washington-shoji/gin-api/helpers"

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

func (handler *ExpHandler) ExpGetAll(ctx *gin.Context) {
	query := `SELECT * from table_exp WHERE deleted_at IS NULL`

	rows, err := handler.Database.Query(query)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Exp not found"}})
		return
	}

	expRes := []*DynamicStruct{}

	for rows.Next() {
		exp, err := scanIntoExp(rows)
		if err != nil {
			fmt.Println("err scan", err)
			return
		}

		expRes = append(expRes, exp)
	}

	fmt.Println("expRes", expRes)

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: expRes})

}

func (handler *ExpHandler) ExpCreate(ctx *gin.Context) {
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

func scanIntoExp(rows *sql.Rows) (*DynamicStruct, error) {
	data := &DynamicStruct{}
	var jsonColumnData []byte
	err := rows.Scan(
		&data.ID,
		&jsonColumnData,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
	)

	if err := json.Unmarshal(jsonColumnData, &data); err != nil {
		fmt.Println("Error unmarshalling JSONB data:", err)
		return nil, err
	}

	return data, err
}
