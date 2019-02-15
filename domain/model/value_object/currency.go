package value_object

// 通貨
type Currency string

// リスト
const (
	USD = Currency("USD")
	EUR = Currency("EUR")
	JPY = Currency("JPY")
)

func (c Currency) New(a int) Money {
	return Money{a, c}
}