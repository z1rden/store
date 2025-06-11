package cart_api

import (
	"cart/internal/cart/logger"
	"cart/internal/cart/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gookit/validate"
	"net/http"
	"strconv"
)

func (a *api) GetCartByUserID() func(http.ResponseWriter, *http.Request) {
	const operation = "api.GetCartByUserID"

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req, err := toGetCartByUserIDRequest(ctx, r)
		if err != nil {
			logger.Errorf(ctx, "%s: userID is not valid: %v", operation, err)
			http.Error(w, fmt.Sprintf("user ID is not valid: %s", err), http.StatusBadRequest)

			return
		}

		cart, err := a.cartService.GetCartByUserID(ctx, req.UserID)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				http.Error(w, fmt.Sprintf("cart for user %d not found", req.UserID), http.StatusNotFound)

				return
			} else {
				logger.Errorf(ctx, "%s: failed to get cart: %v", operation, err)
				http.Error(w, "internal error", http.StatusInternalServerError)

				return
			}
		}

		jsonCart, err := json.Marshal(cart)
		if err != nil {
			logger.Errorf(ctx, "%s: failed to marshal cart: %v", operation, err)
			http.Error(w, fmt.Sprintf("failed to marshal cart: %s", err), http.StatusInternalServerError)

			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err = w.Write(jsonCart)
		if err != nil {
			logger.Errorf(ctx, "%s: failed to write response: %v", operation, err)
			http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)

			return
		}
	}
}

func toGetCartByUserIDRequest(ctx context.Context, r *http.Request) (*GetCartByUserIDRequest, error) {
	const operation = "api.toGetCartByUserIDRequest"

	userID, err := strconv.ParseInt(r.PathValue("user_id"), 10, 64)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to parse user_id: %v", operation, err)

		return nil, err
	}

	req := &GetCartByUserIDRequest{
		UserID: userID,
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return nil, v.Errors
	}

	return req, nil
}
