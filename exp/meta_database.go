package exp

import (
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
)

type MetaDatabaseImp struct {
	Db *sql.DB
}

type Tag struct {
	Table_name string `db:"table_name" json:"table_name"`
}

func NewMetaDatabaseImp(Db *sql.DB) *MetaDatabaseImp {
	return &MetaDatabaseImp{
		Db: Db,
	}
}

func (m *MetaDatabaseImp) MetaDatabase() {

	sqlQuery := `SELECT table_name
	FROM information_schema.tables
	WHERE table_schema = $1;`
	schValue := "public"

	result, err := m.Db.Query(sqlQuery, schValue)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("Showing Tables", result)
	for result.Next() {
		tag := &Tag{}
		err = result.Scan(&tag.Table_name)
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println("Table_name", tag.Table_name)
	}
}
