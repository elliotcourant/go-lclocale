//go:build !wasm

package locale

import (
	"errors"
	"sync"
)

type Currency struct {
	// First three characters are a currency symbol from ISO4217. Fourth character
	// is the separator. Fifth character is '\0'.
	CurrencyCode []byte
	// Local currency symbol.
	CurrencySymbol []byte
	// Radix character.
	MonDecimalPoint []byte
	// Like ThousandsSep above.
	MonThousandsSep []byte
	// Like Grouping above.
	MonGrouping []uint8
}

var (
	currencyCache      map[string]Currency
	currencyCacheMutex sync.RWMutex
)

var (
	ErrCurrencyNotSupported = errors.New("currency not supported")
)

// func GetCurrency(currencyCode string) (*Currency, error) {
//
// }
