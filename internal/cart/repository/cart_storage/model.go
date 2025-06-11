package cart_storage

type Cart struct {
	Items map[int64]uint16 // Map: ключ - итем, uint16 - количество.
}
