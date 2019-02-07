package Controllers

import (
	"apirest/Entities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type paintersControllerStruct struct {
}

func (c paintersControllerStruct) index(w http.ResponseWriter, h *http.Request) {
	fmt.Println("resource painters index")

	painters := Entities.PainterEntity.All()

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(painters)
}

func (c paintersControllerStruct) store(w http.ResponseWriter, h *http.Request) {
	var data Entities.PainterStructure

	decoder := json.NewDecoder(h.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	Entities.PainterEntity.Create(data.Name, data.CityOfOrigin)

	fmt.Println("resource post:" + data.Name)

}

var PaintersController paintersControllerStruct
