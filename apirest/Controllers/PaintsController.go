package Controllers

import (
	"apirest/Entities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type paintsControllerStruct struct {
}

func (c paintsControllerStruct) index(w http.ResponseWriter, h *http.Request) {
	fmt.Println("resource paints index")
	paints := Entities.PaintEntity.All()

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(paints)

}

func (c paintsControllerStruct) store(w http.ResponseWriter, h *http.Request) {
	fmt.Println("resource paints store")

	var data Entities.PaintStructure

	decoder := json.NewDecoder(h.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	Entities.PaintEntity.Create(data.Name, data.CurrentLocation, data.PainterId)

	fmt.Println("resource post:" + data.Name)

}

var PaintsController paintsControllerStruct
