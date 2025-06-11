package stock_api

import (
	"buf.build/go/protovalidate"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"loms/internal/loms/model"
	"loms/pkg/api/stock"
)

func (a *api) Info(ctx context.Context, r *stock.StockInfoRequest) (*stock.StockInfoResponse, error) {
	if err := protovalidate.Validate(r); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	count, err := a.stockService.Info(ctx, r.GetSku())
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("sku %d not found", r.GetSku()))
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("internal error: %v", err))
		}
	}

	return &stock.StockInfoResponse{
		Count: uint64(count),
	}, nil
}
