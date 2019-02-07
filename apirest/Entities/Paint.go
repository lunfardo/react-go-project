package Entities

import (
	"database/sql"
	"log"
)

type PaintEntityStruct struct {
	dbConn *sql.DB
}

type PaintStructure struct {
	Id              int
	PainterId       int
	Name            string
	CurrentLocation string
}

var PaintEntity PaintEntityStruct

func (e *PaintEntityStruct) AddDB(conn *sql.DB) {
	e.dbConn = conn
}

func (e *PaintEntityStruct) All() []PaintStructure {
	var allPaints []PaintStructure
	sqlIndexPainters := `SELECT * FROM paints`
	//Get all Paints
	rows, err := e.dbConn.Query(sqlIndexPainters)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//rows -> []PaintStructure
	for rows.Next() {
		var temp PaintStructure
		err := rows.Scan(&temp.Id, &temp.PainterId, &temp.Name, &temp.CurrentLocation)
		if err != nil {
			log.Fatal(err)
		}
		allPaints = append(allPaints, temp)
	}
	return allPaints
}

func (e *PaintEntityStruct) Create(name string, currentlocation string, idPainter int) {
	sqlStorePaint := `INSERT INTO paints(PainterId,Name,CurrentLocation) VALUES ($1,$2,$3)`
	_, err := e.dbConn.Exec(sqlStorePaint, idPainter, name, currentlocation)
	if err != nil {
		log.Fatal(err)
	}
}
