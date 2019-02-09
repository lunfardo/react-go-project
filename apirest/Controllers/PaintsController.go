package Controllers

import (
	"apirest/Entities"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type paintsControllerStruct struct {
}

func (c paintsControllerStruct) index(w http.ResponseWriter, h *http.Request) {
	fmt.Println("resource paints index")
	paints := Entities.PaintEntity.All().BytesToBase64()

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(paints)

}

func (c paintsControllerStruct) store(w http.ResponseWriter, h *http.Request) {
	fmt.Println("resource paints store")

	//parse multipart form post request
	err := h.ParseMultipartForm(0)
	if err != nil {
		fmt.Println(err)
	}

	//get painter id
	painterId, _ := strconv.Atoi(h.FormValue("PainterId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//make a new painter entity (no stored in db yet) and delete from the struct at the end
	newPaint := Entities.PaintEntity.Make(h.FormValue("Name"), h.FormValue("CurrentLocation"), painterId)
	defer Entities.PaintEntity.Release()

	//Handle file upload https://www.youtube.com/watch?v=lKXkgzEmIUk
	file, _, err := h.FormFile("paintFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	//add the recivied file image to the entity
	err = newPaint.AddImageFile(file)

	//check if file was copied correctly
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = Entities.PaintEntity.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println("resource post:" + newPaint.Name)
}

var PaintsController paintsControllerStruct
