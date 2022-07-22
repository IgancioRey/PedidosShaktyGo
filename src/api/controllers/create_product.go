package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/entities"
	"github.com/IgancioRey/PedidosShaktyGo/src/api/core/usecases"
)

type CreateProduct struct {
	UseCase usecases.CreateProduct
}

func (controller *CreateProduct) Execute(w http.ResponseWriter, r *http.Request) {
	var CreateProductRequest entities.CreateProductRequest
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bodyBytes, &CreateProductRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !CreateProductRequest.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.UseCase.Execute(CreateProductRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
