package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/entities"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/usecases"
)

type CreateOrder struct {
	UseCase usecases.CreateOrder
}

func (controller *CreateOrder) Execute(w http.ResponseWriter, r *http.Request) {
	var CreateOrderRequest entities.CreateOrderRequest
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bodyBytes, &CreateOrderRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !CreateOrderRequest.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var newOrder entities.CreateOrderResponse
	newOrder, err = controller.UseCase.Execute(CreateOrderRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	orderJson, err := json.Marshal(newOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(orderJson)
}
