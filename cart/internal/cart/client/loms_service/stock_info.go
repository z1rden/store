package loms_service

import "context"

func (c *client) StockInfo(ctx context.Context, SkuID int64) (uint16, error) {
	response, err := c.stockGrpcClient.Info(ctx, toStockInfoRequest(SkuID))
	if err != nil {
		return 0, err
	}

	return uint16(response.Count), nil
}
