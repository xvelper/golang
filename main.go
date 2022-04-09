package main

import (
	"database/sql"

	"encoding/json"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type list struct {
	id   int    `json:"id"`
	info string `json:"info"`
}

func api(c *gin.Context) {
	db, err := sql.Open("mysql", "root:root@/test")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.Exec("insert into test.list (info) value (?)", "POPA")

	c.JSON(200, gin.H{
		"message": "",
	})
}
func all(c *gin.Context) {
	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3306)/test")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("select * from test.list")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	lists := []list{}

	for rows.Next() {
		l := list{}
		rows.Scan(&l.id, &l.info)
		lists = append(lists, l)
	}
	b, _ := json.Marshal(lists)

	c.JSON(200, string(b))
}

func main() {
	r := gin.Default()
	r.GET("/api", api)
	r.GET("/all", all)
	r.Run("0.0.0.0:9090")
}
