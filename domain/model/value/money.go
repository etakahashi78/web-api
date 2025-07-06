package value

import (
	"math/big"
)

type Money struct {
	amount *big.Rat
}

func NewMoney(value float64) *Money {
	amount := new(big.Rat).SetFloat64(value)
	return &Money{amount: amount}
}

func (m *Money) Amount() float64 {
	amount, _ := m.amount.Float64()
	return amount
}

func (m *Money) String() string {
	return m.amount.FloatString(2) // 小数点以下2桁で表示
}

func (m *Money) Equals(other *Money) bool {
	return m.amount.Cmp(other.amount) == 0
}

func (m *Money) Add(other *Money) *Money {
	sum := new(big.Rat).Add(m.amount, other.amount)
	return &Money{amount: sum}
}

func (m *Money) Subtract(other *Money) *Money {
	diff := new(big.Rat).Sub(m.amount, other.amount)
	return &Money{amount: diff}
}
