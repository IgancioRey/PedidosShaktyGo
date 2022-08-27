package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/usecases"
	"github.com/go-chi/chi/v5"
)

type GetCustomer struct {
	UseCase usecases.GetCustomer
}

func (controller *GetCustomer) Execute(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	customer, err := controller.UseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	customerJson, err := json.Marshal(customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(customerJson)
}
