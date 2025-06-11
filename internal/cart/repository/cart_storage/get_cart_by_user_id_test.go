package cart_storage_test

import (
	"cart/internal/cart/model"
	"cart/internal/cart/repository/cart_storage"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageGetCartByUserId(t *testing.T) {
	ctx := context.Background()
	s := cart_storage.NewStorage()

	cart, err := s.GetCartByUserID(ctx, 1)
	assert.ErrorIs(t, err, model.ErrNotFound, "Должна вернуться ошибка, так как корзины "+
		"для пользователя еще не существует")
	assert.Nil(t, cart, "Корзина у пользователя еще не существует")

	_ = s.AddItem(ctx, 1, 1, 1)
	cart, err = s.GetCartByUserID(ctx, 1)
	assert.NoError(t, err, "Ошибки не должно возникать, так как у пользователя появилась корзина "+
		"с товаром внутри")
	assert.Len(t, cart.Items, 1, "Пользователь положил один товар нового типа, длина должна быть "+
		"равна 1")
	quantity := cart.Items[1]
	assert.Equal(t, quantity, uint16(1), "Пользователь положил только один товар данного тип, поэтому "+
		"количество должно быть равно 1")

	_ = s.AddItem(ctx, 1, 2, 3)
	cart, err = s.GetCartByUserID(ctx, 1)
	assert.NoError(t, err, "Ошибки не должно возникать, так как у пользователя есть корзина "+
		"с товаром внутри")
	assert.Len(t, cart.Items, 2, "Пользователь положил товар нового типа, длина должна быть "+
		"равна 2")
	quantity = cart.Items[2]
	assert.Equal(t, quantity, uint16(3), "Количество должно быть равным трем")

	_ = s.DeleteItem(ctx, 1, 1)
	cart, err = s.GetCartByUserID(ctx, 1)
	assert.NoError(t, err, "Ошибки не должно возникать, так как у пользователя есть корзина "+
		"с товаром внутри")
	assert.Len(t, cart.Items, 1, "Пользователь удалил только один товар нового типа, длина должна быть "+
		"равна 1")
	quantity = cart.Items[2]
	assert.Equal(t, quantity, uint16(3), "Количество должно быть равным трем")

	_ = s.AddItem(ctx, 1, 1, 1)
	_ = s.DeleteCartByUserID(ctx, 1)
	cart, err = s.GetCartByUserID(ctx, 1)
	assert.ErrorIs(t, err, model.ErrNotFound, "Должна вернуться ошибка, так как корзины "+
		"у пользователя не существует")
	assert.Nil(t, cart, "Корзины после удаления у пользователя не существует")

	_ = s.AddItem(ctx, 1, 2, 3)
	_ = s.DeleteItem(ctx, 1, 2)
	cart, err = s.GetCartByUserID(ctx, 1)
	assert.ErrorIs(t, err, model.ErrNotFound, "Должна вернуться ошибка, так как корзины "+
		"у пользователя не существует после удаления единственного продукта")
	assert.Nil(t, cart, "Корзины после удаления у пользователя не существует")
}
