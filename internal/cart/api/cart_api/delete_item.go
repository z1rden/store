package cart_api

import (
	"cart/internal/cart/logger"
	"context"
	"fmt"
	"github.com/gookit/validate"
	"net/http"
	"strconv"
)

func (a *api) DeleteItem() func(w http.ResponseWriter, r *http.Request) {
	const operation = "api.DeleteItem"

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		req, err := toDeleteItemRequest(ctx, r)
		if err != nil {
			logger.Errorf(ctx, "%s: request is not valid: %v", operation, err)
			http.Error(w, fmt.Sprintf("request is not valid %s", err), http.StatusBadRequest)

			return
		}

		err = a.cartService.DeleteItem(ctx, req.UserID, req.SkuID)
		if err != nil {
			logger.Errorf(ctx, "%s: failed to delete item: %v", operation, err)
			http.Error(w, fmt.Sprintf("failed to delete item: %s", err), http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func toDeleteItemRequest(ctx context.Context, r *http.Request) (*DeleteItemRequest, error) {
	const operation = "api.toDeleteItemRequest"

	userID, err := strconv.ParseInt(r.PathValue("user_id"), 10, 64)
	if err != nil {
		logger.Errorf(ctx, "%s: userID is not valid: %v", operation, err)

		return nil, err
	}

	skuID, err := strconv.ParseInt(r.PathValue("sku_id"), 10, 64)
	if err != nil {
		logger.Errorf(ctx, "%s: skuID is not valid: %v", operation, err)

		return nil, err
	}

	req := &DeleteItemRequest{
		UserID: userID,
		SkuID:  skuID,
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return nil, v.Errors
	}

	return req, nil
}
