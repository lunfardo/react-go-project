package Entities

import (
	"database/sql"
	"log"
)

type PainterEntityStruct struct {
	dbConn *sql.DB
}

type PainterStructure struct {
	Id           int
	Name         string
	CityOfOrigin string
	Diedrich     bool
}

var PainterEntity PainterEntityStruct

func (e *PainterEntityStruct) AddDB(conn *sql.DB) {
	e.dbConn = conn
}

func (e *PainterEntityStruct) All() []PainterStructure {
	var allPainters []PainterStructure
	sqlIndexPainters := `SELECT * FROM painters`
	rows, err := e.dbConn.Query(sqlIndexPainters)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var temp PainterStructure
		err := rows.Scan(&temp.Id, &temp.Name, &temp.CityOfOrigin, &temp.Diedrich)
		if err != nil {
			log.Fatal(err)
		}
		allPainters = append(allPainters, temp)
	}
	return allPainters
}

func (e *PainterEntityStruct) Create(name string, cityoforigin string) {
	sqlStorePainter := `INSERT INTO painters(Fullname,CityOfOrigin) VALUES ($1,$2)`
	_, err := e.dbConn.Exec(sqlStorePainter, name, cityoforigin)
	if err != nil {
		log.Fatal(err)
	}
}
