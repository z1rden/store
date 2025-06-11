package cart_api

import (
	"cart/internal/cart/model"
	"cart/internal/cart/service/cart_service"
	"fmt"
	"net/http"
)

type API interface {
	GetHandlers() []model.HttpAPIHandler
	AddItem() func(http.ResponseWriter, *http.Request)
	DeleteItem() func(http.ResponseWriter, *http.Request)
	DeleteCartByUserID() func(http.ResponseWriter, *http.Request)
	GetCartByUserID() func(http.ResponseWriter, *http.Request)
	Checkout() func(http.ResponseWriter, *http.Request)
}

type api struct {
	cartService cart_service.Service
}

func NewApi(cartService cart_service.Service) API {
	return &api{
		cartService: cartService,
	}
}

func (a *api) GetHandlers() []model.HttpAPIHandler {
	return []model.HttpAPIHandler{
		{
			Pattern: fmt.Sprintf("%s /user/{%s}/cart/{%s}", http.MethodPost, "user_id", "sku_id"),
			Handler: a.AddItem(),
		},
		{
			Pattern: fmt.Sprintf("%s /user/{%s}/cart/{%s}", http.MethodDelete, "user_id", "sku_id"),
			Handler: a.DeleteItem(),
		},
		{
			Pattern: fmt.Sprintf("%s /user/{%s}/cart", http.MethodDelete, "user_id"),
			Handler: a.DeleteCartByUserID(),
		},
		{
			Pattern: fmt.Sprintf("%s /user/{%s}/cart", http.MethodGet, "user_id"),
			Handler: a.GetCartByUserID(),
		},
		{
			Pattern: fmt.Sprintf("%s /cart/checkout", http.MethodPost),
			Handler: a.Checkout(),
		},
	}
}
