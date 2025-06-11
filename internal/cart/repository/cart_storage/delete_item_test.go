package cart_storage_test

import (
	"cart/internal/cart/repository/cart_storage"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageDeleteItem(t *testing.T) {
	ctx := context.Background()
	s := cart_storage.NewStorage()

	err := s.DeleteItem(ctx, 1, 1)
	assert.NoError(t, err, "Ошибки не должно возникать при удалении продукта из пустой корзины")

	err = s.AddItem(ctx, 1, 1, 2)
	err = s.DeleteItem(ctx, 1, 1)
	assert.NoError(t, err, "Ошибки не должно возникать при удалении продукта из непустой корзины")

	cart, err := s.GetCartByUserID(ctx, 1)
	assert.Nil(t, cart, "Корзины у пользователя не должно быть пустой после удаления единственного товара")
}
