package Entities

import (
	"database/sql"
	"encoding/base64"
	"io/ioutil"
	"log"
	"mime/multipart"
)

type PaintEntityStruct struct {
	dbConn *sql.DB
	entity *PaintStructure
}

type PaintStructure struct {
	Id              int
	PainterId       int
	Name            string
	CurrentLocation string
	Image           []byte
}

type PaintStructureBASE64 struct {
	Id              int
	PainterId       int
	Name            string
	CurrentLocation string
	Image           string
}

type PaintStructureArray []PaintStructure
type PaintStructureBASE64Array []PaintStructureBASE64

var PaintEntity PaintEntityStruct

func (s *PaintEntityStruct) AddDB(conn *sql.DB) {
	s.dbConn = conn
}

func (s *PaintEntityStruct) All() PaintStructureArray {
	var allPaints PaintStructureArray
	sqlIndexPainters := `SELECT * FROM paints`
	//Get all Paints
	rows, err := s.dbConn.Query(sqlIndexPainters)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//rows -> []PaintStructure
	for rows.Next() {
		var temp PaintStructure
		err := rows.Scan(&temp.Id, &temp.PainterId, &temp.Name, &temp.CurrentLocation, &temp.Image)
		if err != nil {
			log.Fatal(err)
		}

		allPaints = append(allPaints, temp)
	}
	return allPaints
}

func (s *PaintEntityStruct) Make(name string, currentlocation string, idPainter int) *PaintStructure {
	s.entity = &PaintStructure{Name: name, CurrentLocation: currentlocation, PainterId: idPainter}
	return s.entity
}

func (e *PaintStructure) AddImageFile(f multipart.File) error {
	bytes, err := ioutil.ReadAll(f)
	e.Image = bytes
	return err
}

func (s *PaintEntityStruct) Save() error {
	sqlStorePaint := `INSERT INTO paints(PainterId,Name,CurrentLocation,Image) VALUES ($1,$2,$3, $4)`
	_, err := s.dbConn.Exec(sqlStorePaint, s.entity.PainterId, s.entity.Name, s.entity.CurrentLocation, s.entity.Image)
	return err
}

func (s *PaintEntityStruct) Release() {
	s.entity = nil
}

func (a PaintStructureArray) BytesToBase64() PaintStructureBASE64Array {

	result := make(PaintStructureBASE64Array, len(a))
	for index, element := range a {
		result[index] = PaintStructureBASE64{
			Id:              element.Id,
			PainterId:       element.PainterId,
			CurrentLocation: element.CurrentLocation,
			Image:           base64.StdEncoding.EncodeToString(element.Image),
		}
	}
	return result
}
