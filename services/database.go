package service

import (
	"fmt"
	"log"
	
	"github.com/go-pg/pg/v10"
)

func PgDataBase() (con *pg.DB) {
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