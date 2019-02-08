package Entities

import (
	"database/sql"
	"encoding/base64"
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
	Image           []byte
}

type PaintStructureJSON struct {
	Id              int
	PainterId       int
	Name            string
	CurrentLocation string
	Image           string
}

var PaintEntity PaintEntityStruct

func (e *PaintEntityStruct) AddDB(conn *sql.DB) {
	e.dbConn = conn
}

func (e *PaintEntityStruct) All() []PaintStructureJSON {
	var allPaints []PaintStructureJSON
	sqlIndexPainters := `SELECT * FROM paints`
	//Get all Paints
	rows, err := e.dbConn.Query(sqlIndexPainters)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//rows -> []PaintStructure
	for rows.Next() {
		var saveByteArray []byte
		var temp PaintStructureJSON
		err := rows.Scan(&temp.Id, &temp.PainterId, &temp.Name, &temp.CurrentLocation, &saveByteArray)
		if err != nil {
			log.Fatal(err)
		}

		temp.Image = base64.StdEncoding.EncodeToString(saveByteArray)

		allPaints = append(allPaints, temp)
	}
	return allPaints
}

func (e *PaintEntityStruct) Create(name string, currentlocation string, idPainter int, bytes []byte) {
	sqlStorePaint := `INSERT INTO paints(PainterId,Name,CurrentLocation,Image) VALUES ($1,$2,$3, $4)`
	_, err := e.dbConn.Exec(sqlStorePaint, idPainter, name, currentlocation, bytes)
	if err != nil {
		log.Fatal(err)
	}
}
