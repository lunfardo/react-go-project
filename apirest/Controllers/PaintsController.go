package Controllers

import (
	"apirest/Entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

	// Calling ParseMultipartForm, with maxMemory = 0. Tried also with 32, 10000 etc. but without success.
	err := h.ParseMultipartForm(0)
	if err != nil {
		fmt.Println(err)
	}

	//Handle file upload https://www.youtube.com/watch?v=lKXkgzEmIUk
	file, headers, err := h.FormFile("paintFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	fmt.Println("|nfile:", file, "|nheader:", headers, "|nerr", err)

	//read the file
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	s := string(bytes)
	fmt.Println(s)
	/**/

	painterId, _ := strconv.Atoi(h.FormValue("PainterId"))

	Entities.PaintEntity.Create(h.FormValue("Name"), h.FormValue("CurrentLocation"), painterId, bytes)

	fmt.Println("resource post:" + data.Name)

}

var PaintsController paintsControllerStruct
