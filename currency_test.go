package locale

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrencyMapping(t *testing.T) {
	j, _ := json.MarshalIndent(currencyMapping, "", "  ")
	fmt.Println(string(j))
}

func TestFractionalDigits(t *testing.T) {
	t.Run("USD", func(t *testing.T) {
		fraction, err := GetCurrencyInternationalFractionalDigits("USD")
		assert.NoError(t, err, "should retrieve fractional digits for USD")
		assert.EqualValues(t, 2, fraction)
	})

	t.Run("JPY", func(t *testing.T) {
		fraction, err := GetCurrencyInternationalFractionalDigits("JPY")
		assert.NoError(t, err, "should retrieve fractional digits for JPY")
		assert.EqualValues(t, 0, fraction)
	})
}

func TestGetInstalledCurrencies(t *testing.T) {
	t.Run("USD", func(t *testing.T) {
		installed := GetInstalledCurrencies()
		assert.Contains(t, installed, "USD", "must contain USD currency")
	})

	t.Run("EUR", func(t *testing.T) {
		installed := GetInstalledCurrencies()
		assert.Contains(t, installed, "EUR", "must contain EUR currency")
	})

	t.Run("JPY", func(t *testing.T) {
		installed := GetInstalledCurrencies()
		assert.Contains(t, installed, "JPY", "must contain JPY currency")
	})
}
