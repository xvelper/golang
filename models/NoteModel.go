package model

import (
	service "notice/services"
)

type Note struct {
	tableName struct{} `pg:"notes"`
	Id        int64    `json:"id" pg:"id,pk"`
	CreatedAt string   `json:"created_at", pg:"created_at"`
	Title     string   `json:"title", pg:"title"`
	Info      string   `json:"info", pg:"info"`
}

func SelectNotes() []Note {
	var notes []Note
	db := service.PgDataBase()

	err := db.Model(&notes).Select()

	if err != nil {
		panic(err)
	}

	db.Close()

	return notes
}

func SelectNote(id int64) Note {
	var note Note
	db := service.PgDataBase()

	err := db.Model(&note).Where("id = ?", id).Select()

	if err != nil {
		panic(err)
	}

	db.Close()

	return note
}

func InsertNote(note Note) Note {
	db := service.PgDataBase()

	_, err := db.Model(&note).Insert()
	if err != nil {
		panic(err)
	}

	db.Close()

	return note
}

func UpdateNote(note Note) Note {
	db := service.PgDataBase()

	_, err := db.Model(&note).Where("id = ?", note.Id).Update()
	if err != nil {
		panic(err)
	}

	db.Close()

	return note
}

func DeleteNote(id int64) error {
	var note Note
	db := service.PgDataBase()

	_, err := db.Model(&note).Where("id = ?", id).Delete()

	db.Close()
	return err
}
