package Controllers

import "net/http"

/*type Controller struct {
	Dbconn *sql.DB
}*/

type Resource interface {
	index(w http.ResponseWriter, h *http.Request)
	store(http.ResponseWriter, *http.Request)
}

type Controller struct {
	ResourceController interface {
		index(w http.ResponseWriter, h *http.Request)
		store(http.ResponseWriter, *http.Request)
	}
}

func (c *Controller) ResourceMux(w http.ResponseWriter, h *http.Request) {
	switch h.Method {
	case http.MethodGet:
		c.ResourceController.index(w, h)

	case http.MethodPost:
		c.ResourceController.store(w, h)
	}

}

//api/v1/painters
func (c *Controller) AddRoutesToMux(path string, m *http.ServeMux) {
	m.HandleFunc(path, c.ResourceMux)
}
