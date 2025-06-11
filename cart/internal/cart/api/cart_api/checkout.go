package cart_api

import (
	"cart/internal/cart/logger"
	"cart/internal/cart/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gookit/validate"
	"io"
	"net/http"
)

func (a *api) Checkout() func(http.ResponseWriter, *http.Request) {
	const operation = "api.Checkout"

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		req, err := toCheckoutRequest(r)
		if err != nil {
			logger.Errorf(ctx, "%s: request is not valid: %v", operation, err)
			http.Error(w, fmt.Sprintf("request is not valid: %s", err), http.StatusBadRequest)

			return
		}

		orderID, err := a.cartService.Checkout(ctx, req.UserID)
		if err != nil {
			logger.Errorf(ctx, "%s: failed to checkout: %v", operation, err)
			if errors.Is(err, model.ErrNotFound) {
				http.Error(w, "cart not found", http.StatusNotFound)
			} else {
				http.Error(w, "internal error", http.StatusInternalServerError)
			}

			return
		}

		err = toCheckoutResponse(ctx, w, orderID)
		if err != nil {
			logger.Errorf(ctx, "%s :failed to write response: %v", operation, err)
		}
	}
}

func toCheckoutRequest(r *http.Request) (*CheckoutRequest, error) {
	const operation = "api.toCheckoutRequest"

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	req := &CheckoutRequest{}
	err = json.Unmarshal(body, req)
	if err != nil {
		logger.Errorf(r.Context(), "%s: failed to unmarshal request body: %v", operation, err)
		return nil, err
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return nil, v.Errors
	}

	return req, nil
}

func toCheckoutResponse(ctx context.Context, w http.ResponseWriter, orderID int64) error {
	const operation = "api.toCheckoutResponse"

	resp := &CheckoutResponse{
		OrderID: orderID,
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to marshal response: %v", operation, err)
		http.Error(w, "internal error", http.StatusInternalServerError)

		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(respJson)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to write response: %v", operation, err)
		http.Error(w, "internal error", http.StatusInternalServerError)

		return err
	}

	return nil
}
