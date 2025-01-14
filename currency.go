//go:build !wasm

package locale

import (
	"errors"
	"strings"
)

var (
	ErrCurrencyNotSupported = errors.New("currency not supported")
)

func GetCurrencyInternationalFractionalDigits(currency string) (int64, error) {
	currency = strings.ToUpper(currency)
	locales, ok := currencyMapping[currency]
	if !ok {
		return -1, ErrCurrencyNotSupported
	}

	lconv, err := GetLConv(locales[0])
	if err != nil {
		return -1, err
	}

	return int64(lconv.IntFracDigits), nil
}
