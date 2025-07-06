package presentors

import (
	"time"
)

type ProductPresenter struct{}

func NewProductPresenter() *ProductPresenter {
	return &ProductPresenter{}
}

// ProductResponse は単一の商品を JSON で返すためのレスポンス構造体です。
type ProductResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//type ProductsResponse []ProductResponse
//
//// PresentProduct は model.Product を ProductResponse に変換します。
//func (p *ProductPresenter) PresentProduct(product *model.Product) ProductResponse {
//	return ProductResponse{
//		ID:        product.ID,
//		Name:      product.Name,
//		Price:     product.Price,
//		CreatedAt: product.CreatedAt,
//		UpdatedAt: product.UpdatedAt,
//	}
//}
//
//// PresentProducts は []*model.Product を ProductsResponse に変換します。
//func (p *ProductPresenter) PresentProducts(products []*model.Product) ProductsResponse {
//	response := make(ProductsResponse, len(products))
//	for i, product := range products {
//		response[i] = p.PresentProduct(product)
//	}
//	return response
//}
//
//// RespondProduct は単一の商品を JSON レスポンスとして書き込みます。
//func (p *ProductPresenter) RespondProduct(w http.ResponseWriter, statusCode int, product *model.Product) {
//	response := p.PresentProduct(product)
//	w.WriteHeader(statusCode)
//	w.Header().Set("Content-Type", "application/json; charset=utf-8")
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		// エラーハンドリングは必要に応じて
//		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
//		return
//	}
//}
//
//// RespondProducts は複数の商品を JSON レスポンスとして書き込みます。
//func (p *ProductPresenter) RespondProducts(w http.ResponseWriter, statusCode int, products []*model.Product) {
//	response := p.PresentProducts(products)
//	w.WriteHeader(statusCode)
//	w.Header().Set("Content-Type", "application/json; charset=utf-8")
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		// エラーハンドリングは必要に応じて
//		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
//		return
//	}
//}
