package Entities

import (
	"database/sql"
	"encoding/base64"
	"io/ioutil"
	"log"
	"mime/multipart"
)

type PainterEntityStruct struct {
	dbConn *sql.DB
	entity *PainterStructure
}

type PainterStructure struct {
	Id           int
	Name         string
	CityOfOrigin string
	Diedrich     bool
	Image        []byte
}

type PainterStructureBASE64 struct {
	Id           int
	Name         string
	CityOfOrigin string
	Diedrich     bool
	Image        string
}

type PainterStructureArray []PainterStructure
type PainterStructureBASE64Array []PainterStructureBASE64

var PainterEntity PainterEntityStruct

func (s *PainterEntityStruct) AddDB(conn *sql.DB) {
	s.dbConn = conn
}

func (s *PainterEntityStruct) All() PainterStructureArray {
	var allPainters PainterStructureArray

	err := s.dbConn.Ping()
	if err != nil {
		return allPainters
	}

	sqlIndexPainters := `SELECT * FROM painters`
	rows, err := s.dbConn.Query(sqlIndexPainters)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var temp PainterStructure
		err := rows.Scan(&temp.Id, &temp.Name, &temp.CityOfOrigin, &temp.Diedrich, &temp.Image)
		if err != nil {
			log.Printf(err.Error())
		}
		allPainters = append(allPainters, temp)
	}
	return allPainters
}

func (s *PainterEntityStruct) Make(name string, cityoforigin string) *PainterStructure {
	s.entity = &PainterStructure{Name: name, CityOfOrigin: cityoforigin}
	return s.entity
}

func (e *PainterStructure) AddImageFile(f multipart.File) error {
	bytes, err := ioutil.ReadAll(f)
	e.Image = bytes
	return err
}

func (s *PainterEntityStruct) Save() error {
	sqlStorePainter := `INSERT INTO painters(Fullname,CityOfOrigin,Image) VALUES ($1,$2,$3)`
	_, err := s.dbConn.Exec(sqlStorePainter, s.entity.Name, s.entity.CityOfOrigin, s.entity.Image)
	return err
}

func (s *PainterEntityStruct) Release() {
	s.entity = nil
}

func (a PainterStructureArray) BytesToBase64() PainterStructureBASE64Array {

	result := make(PainterStructureBASE64Array, len(a))
	for index, element := range a {
		result[index] = PainterStructureBASE64{
			Id:           element.Id,
			Name:         element.Name,
			CityOfOrigin: element.CityOfOrigin,
			Diedrich:     element.Diedrich,
			Image:        base64.StdEncoding.EncodeToString(element.Image),
		}
	}
	return result
}
