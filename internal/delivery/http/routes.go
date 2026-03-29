package http

import "net/http"

func MapRoutes(mux *http.ServeMux, h *ProductHandler) {
	mux.HandleFunc("GET /products", h.GetAll)
	mux.HandleFunc("GET /product/{id}", h.GetById)
	mux.HandleFunc("POST /product", h.Create)
	mux.HandleFunc("PATCH /product/{id}", h.Update)
	mux.HandleFunc("DELETE /product/{id}", h.Delete)
}
