package controller

import (
	"net/http"
	"strconv"

	model "notice/models"

	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.SelectNotes())
}

func GetNote(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	c.IndentedJSON(http.StatusOK, model.SelectNote(id))
}

func DelNote(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := model.DeleteNote(id)

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
	var note model.Note

	err := c.BindJSON(&note)
	if err != nil {
		return
	}

	note.CreatedAt = "2022-04-16T00:00:00"

	c.IndentedJSON(http.StatusOK, model.InsertNote(note))
}

func EditNote(c *gin.Context) {
	var note model.Note

	err := c.BindJSON(&note)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, model.UpdateNote(note))
}
