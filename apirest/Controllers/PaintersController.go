package Controllers

import (
	"apirest/Entities"
	"encoding/json"
	"fmt"
	"net/http"
)

type paintersControllerStruct struct {
}

func (c paintersControllerStruct) index(w http.ResponseWriter, h *http.Request) {
	fmt.Println("resource painters index")

	painters := Entities.PainterEntity.All().BytesToBase64()

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(painters)
}

func (c paintersControllerStruct) store(w http.ResponseWriter, h *http.Request) {
	fmt.Println("resource painters store")

	//parse multipart form post request
	err := h.ParseMultipartForm(0)
	if err != nil {
		fmt.Println(err)
	}

	//make a new painter entity (no stored in db yet) and delete from the struct at the end
	newPainter := Entities.PainterEntity.Make(h.FormValue("Name"), h.FormValue("CityOfOrigin"))
	defer Entities.PainterEntity.Release()

	//get the painter profile picture
	file, _, err := h.FormFile("painterFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	//add the recivied file image to the entity
	err = newPainter.AddImageFile(file)

	//check if file was copied correctly
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//Store the new entity in the DB
	err = Entities.PainterEntity.Save()

	//check the entity was succesfully stored
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println("resource post:" + newPainter.Name)

}

var PaintersController paintersControllerStruct
