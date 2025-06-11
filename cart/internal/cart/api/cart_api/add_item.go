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
	"strconv"
)

func (a *api) AddItem() func(w http.ResponseWriter, r *http.Request) {
	const operation = "api.AddItem"

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		req, err := toAddItemRequest(ctx, r)
		if err != nil {
			logger.Errorf(ctx, "%s: request is not valid: %v", operation, err)
			http.Error(w, fmt.Sprintf("request is not valid: %s", err), http.StatusBadRequest)

			return
		}

		err = a.cartService.AddItem(r.Context(), req.UserID, req.SkuID, req.Quantity)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				http.Error(w, fmt.Sprintf("skuID %d not found: %s", req.SkuID, err), http.StatusNotFound)

				return
			} else {
				logger.Errorf(ctx, "%s: failed to add item: %v", operation, err)
				http.Error(w, fmt.Sprintf("failed to add item: %s", err), http.StatusInternalServerError)

				return
			}
		}

		w.WriteHeader(http.StatusOK)
	}
}

func toAddItemRequest(ctx context.Context, r *http.Request) (*AddItemRequest, error) {
	const operation = "api.toAddItemRequest"

	userID, err := strconv.ParseInt(r.PathValue("user_id"), 10, 64)
	if err != nil {
		logger.Errorf(ctx, "%s: userID is not valid: %v", operation, err)

		return nil, fmt.Errorf("userID is not valid: %s", err)
	}

	skuID, err := strconv.ParseInt(r.PathValue("sku_id"), 10, 64)
	if err != nil {
		logger.Errorf(ctx, "%s: skuID is not valid: %v", operation, err)

		return nil, fmt.Errorf("skuID is not valid: %s", err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to read body: %v", operation, err)

		return nil, err
	}

	data := &AddItemRequestBody{}
	err = json.Unmarshal(body, data)
	if err != nil {
		logger.Errorf(ctx, "%s: failed to unmarshal body json: %v", operation, err)

		return nil, err
	}

	req := &AddItemRequest{
		UserID:   userID,
		SkuID:    skuID,
		Quantity: data.Count,
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return nil, v.Errors
	}

	return req, nil
}
