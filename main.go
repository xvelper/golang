package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type Note struct {
	tableName struct{} `pg:"notes"`
	Id        int64    `json:"id" pg:"id,pk"`
	CreatedAt string   `json:"created_at" pg:"created_at"`
	Title     string   `json:"title" pg:"title"`
	Info      string   `json:"info" pg:"info"`
}

func SelectNotes() []Note {
	var notes []Note
	db := pgDataBase()

	err := db.Model(&notes).Select()

	if err != nil {
		panic(err)
	}

	db.Close()

	return notes
}

func InsertNote(note Note) Note {
	db := pgDataBase()

	_, err := db.Model(&note).Insert()

	if err != nil {
		panic(err)
	}

	db.Close()

	return note
}

func SelectNote(id int64) Note {
	var note Note
	db := pgDataBase()

	err := db.Model(&note).Where("id = ?", id).Select()

	if err != nil {
		panic(err)
	}

	db.Close()

	return note

}

func pgDataBase() *pg.DB {
	address := fmt.Sprintf("%s:%s", "localhost", "5432")
	options := &pg.Options{
		Addr:     address,
		User:     "postgres",
		Password: "postgres",
		Database: "notice",
		PoolSize: 50,
	}

	con := pg.Connect(options)

	if con == nil {
		log.Fatal("Нет подключения к БД")
	}
	return con
}

func Api(c *gin.Context) {
	db := pgDataBase()

	db.Close()

	c.JSON(200, gin.H{
		"api": "notice",
	})
}

func GetNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, SelectNotes())
}

func AddNote(c *gin.Context) {
	var note Note

	err := c.BindJSON(&note)
	if err != nil {
		return
	}

	note.CreatedAt = "2022-04-23T00:00:00"

	c.IndentedJSON(http.StatusOK, InsertNote(note))
}

func GetNote(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	c.IndentedJSON(http.StatusOK, SelectNote(id))
}

func main() {
	r := gin.Default()
	rApi := r.Group("/api")
	{
		r.GET("/notes", GetNotes)
		rNote := rApi.Group("/note")
		{
			rNote.POST("/add", AddNote)
			rNote.GET("/:id", GetNote)
		}
	}

	r.Run("0.0.0.0:9090")
}
