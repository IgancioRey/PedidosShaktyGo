package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/usecases"
)

type GetProducts struct {
	UseCase usecases.GetProducts
}

func (controller *GetProducts) Execute(w http.ResponseWriter, r *http.Request) {
	result, err := controller.UseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(resultJson)

}
