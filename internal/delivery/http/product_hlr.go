package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Azmi117/API-USER2.git/internal/models"
	"github.com/Azmi117/API-USER2.git/internal/pkg/apperror"
	"github.com/Azmi117/API-USER2.git/internal/usecase"
)

type ProductHandler struct {
	usecase *usecase.ProductUsecase
}

func NewProductHandler(params *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		usecase: params,
	}
}

func SendError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	if appErr, ok := err.(*apperror.Apperror); ok {
		w.WriteHeader(appErr.Code)
		json.NewEncoder(w).Encode(appErr)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"message": "Internal Server Error"})
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.usecase.GetAll()

	if err != nil {
		SendError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	res, err := h.usecase.GetById(id)

	if err != nil {
		SendError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Product

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		SendError(w, apperror.Internal("Invalid format json"))
		return
	}

	res, err := h.usecase.Create(input)

	if err != nil {
		SendError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	var input models.Product
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		SendError(w, apperror.Internal("Invalid format json"))
		return
	}

	res, err := h.usecase.Update(id, input)

	if err != nil {
		SendError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	if err := h.usecase.Delete(id); err != nil {
		SendError(w, err)
	}

	w.WriteHeader(http.StatusNoContent)
}
