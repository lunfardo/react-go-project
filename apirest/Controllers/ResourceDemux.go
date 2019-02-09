package Controllers

import "net/http"

type ResourceDemux struct {
	ResourceController interface {
		index(w http.ResponseWriter, h *http.Request)
		store(http.ResponseWriter, *http.Request)
	}
}

func (c *ResourceDemux) ResourceMux(w http.ResponseWriter, h *http.Request) {
	switch h.Method {
	case http.MethodGet:
		c.ResourceController.index(w, h)

	case http.MethodPost:
		c.ResourceController.store(w, h)
	}

}

func (c *ResourceDemux) AddRoutesToMux(path string, m *http.ServeMux) {
	m.HandleFunc(path, c.ResourceMux)
}
