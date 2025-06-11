package cart_api

import (
	"cart/internal/cart/logger"
	"context"
	"fmt"
	"github.com/gookit/validate"
	"net/http"
	"strconv"
)

func (a *api) DeleteCartByUserID() func(w http.ResponseWriter, r *http.Request) {
	const operation = "api.DeleteCartByUserID"

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		req, err := toDeleteCartByUserIdRequest(ctx, r)
		if err != nil {
			logger.Errorf(ctx, "%s: request is not valid: %v", operation, err)
			http.Error(w, fmt.Sprintf("request is not valid: %v", err), http.StatusBadRequest)

			return
		}

		err = a.cartService.DeleteCartByUserId(ctx, req.UserID)
		if err != nil {
			logger.Errorf(ctx, "%s: failed to delete cart by userId: %v", operation, err)
			http.Error(w, fmt.Sprintf("failed to delete cart by userId: %v", err), http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func toDeleteCartByUserIdRequest(ctx context.Context, r *http.Request) (*DeleteItemsByUserIDRequest, error) {
	const operation = "api.toDeleteCartByUserIDRequest"

	userID, err := strconv.ParseInt(r.PathValue("user_id"), 10, 64)
	if err != nil {
		logger.Errorf(ctx, "%s: userID is not valid: %v", operation, err)

		return nil, fmt.Errorf("userID is not valid: %s", err)
	}

	req := &DeleteItemsByUserIDRequest{
		UserID: userID,
	}

	v := validate.Struct(req)
	if !v.Validate() {
		return nil, v.Errors
	}

	return req, nil
}
