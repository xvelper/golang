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
	CreatedAt string   `json:"created_at", pg:"created_at"`
	Title     string   `json:"title", pg:"title"`
	Info      string   `json:"info", pg:"info"`
}

func pgDataBase() (con *pg.DB) {
	address := fmt.Sprintf("%s:%s", "localhost", "5432")
	options := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     address,
		Database: "notice",
	}
	con = pg.Connect(options)
	if con == nil {
		log.Fatal("Нет подключения к БД!")
	}
	return
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

func InsertNote(note Note) Note {
	db := pgDataBase()

	_, err := db.Model(&note).Insert()
	if err != nil {
		panic(err)
	}

	db.Close()

	return note
}

func UpdateNote(note Note) Note {
	db := pgDataBase()

	_, err := db.Model(&note).Where("id = ?", note.Id).Update()
	if err != nil {
		panic(err)
	}

	db.Close()

	return note
}

func DeleteNote(id int64) error {
	db := pgDataBase()
	var note Note

	_, err := db.Model(&note).Where("id = ?", id).Delete()

	db.Close()
	return err
}

func Api(c *gin.Context) {
	c.JSON(200, gin.H{
		"api": "notice",
	})
}

func GetNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, SelectNotes())
}

func GetNote(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	c.IndentedJSON(http.StatusOK, SelectNote(id))
}

func DelNote(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := DeleteNote(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Возникла ошибка при удалении объекта",
		})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Объект успешно удалён",
		})
	}
}

func AddNote(c *gin.Context) {
	var note Note

	err := c.BindJSON(&note)
	if err != nil {
		return
	}

	note.CreatedAt = "2022-04-16T00:00:00"

	c.IndentedJSON(http.StatusOK, InsertNote(note))
}

func EditNote(c *gin.Context) {
	var note Note

	err := c.BindJSON(&note)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, UpdateNote(note))
}

func main() {
	r := gin.Default()
	r.GET("/api", Api)

	rApi := r.Group("/api")
	{
		rApi.GET("/notes", GetNotes)
		rNote := rApi.Group("/note")
		{
			rNote.POST("/add", AddNote)
			rNote.GET("/:id", GetNote)
			rNote.PUT("/edit", EditNote)
			rNote.DELETE("/:id", DelNote)
		}
	}

	r.Run("0.0.0.0:9090")
}
