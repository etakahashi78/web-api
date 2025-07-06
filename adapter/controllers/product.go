package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"web-api/usecase/interactor"
)

// ProductController は商品の HTTP リクエストを処理するコントローラーです。
type ProductController struct {
	productInteractor interactor.ProductInteractor
}

func NewProductController(interactor interactor.ProductInteractor) *ProductController {
	return &ProductController{productInteractor: interactor}
}

// RegisterProductHandler ...
func (pc *ProductController) RegisterProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		ResponseError(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	product, err := pc.productInteractor.RegisterProduct(ctx, requestBody.Name, requestBody.Price)
	if err != nil {
		log.Printf("failed to register product err:%v", err)
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	ResponseJSON(w, http.StatusCreated, product)
}

// ListProductsHandler ...
func (pc *ProductController) ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	products, err := pc.productInteractor.ListProducts(ctx)
	if err != nil {
		log.Printf("failed to list products err:%v", err)
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	ResponseJSON(w, http.StatusOK, products)
}

// GetProductHandler ...
func (pc *ProductController) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		ID int64 `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		ResponseError(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	products, err := pc.productInteractor.GetProduct(ctx, requestBody.ID)
	if err != nil {
		log.Printf("failed to list products err:%v", err)
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	ResponseJSON(w, http.StatusOK, products)
}
