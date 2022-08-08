package stock

type StockNotFoundError struct{}

func (*StockNotFoundError) Error() string {
	return "stock not found"
}
