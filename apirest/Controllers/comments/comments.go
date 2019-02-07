package comments

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type dataForm struct {
	Name string
	Time string
}

func index(w http.ResponseWriter, h *http.Request) {
	fmt.Println("resource comments index")
}

func store(w http.ResponseWriter, h *http.Request) {
	var data dataForm
	decoder := json.NewDecoder(h.Body)
	err := decoder.Decode(&data.Name)
	if err != nil {
		panic(err)
	}
	fmt.Println("resource post:" + data.Name)

}

func ResourceMux(w http.ResponseWriter, h *http.Request) {
	switch h.Method {
	case http.MethodGet:
		index(w, h)

	case http.MethodPost:
		store(w, h)
	}

}

func AddRoutesToMux(m *http.ServeMux) {
	m.HandleFunc("/comments", ResourceMux)
}
