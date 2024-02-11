package locale_test

import (
	"testing"

	locale "github.com/elliotcourant/go-lclocale"
	"github.com/stretchr/testify/assert"
)

func TestGetLConv(t *testing.T) {
	t.Run("en_US", func(t *testing.T) {
		lconv, err := locale.GetLConv("en_US")
		assert.NoError(t, err, "should not return an error for en_US")
		assert.Equal(t, []byte("."), lconv.DecimalPoint)
		assert.Equal(t, []byte(","), lconv.ThousandsSep)
		assert.Equal(t, []byte("$"), lconv.CurrencySymbol)
	})
}
