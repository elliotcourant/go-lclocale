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

// GetInstalledCurrencies returns a list of ISO currency codes that the current
// system has information on. Currency codes not in this list may still be valid
// but the current system has no information on them.
func GetInstalledCurrencies() []string {
	return installedCurrencies
}
