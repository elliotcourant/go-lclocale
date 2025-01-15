//go:build !wasm

package locale

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetLocale(t *testing.T) {
	locale, err := setlocale("bogus")
	assert.Empty(t, locale, "returned locale should be empty")
	assert.EqualError(t, err, "failed to set locale, wanted: [bogus] got: []")

	locale, err = setlocale("C")
	assert.Equal(t, "C", locale, "returned locale should match input")
	assert.NoError(t, err, "should be able to set locale to C")

	locale, err = setlocale("")
	assert.Equal(t, "C", locale, "passing blank should return the current locale")
	assert.NoError(t, err, "should not return an error")
}

func TestLocaleConv(t *testing.T) {
	t.Run("C", func(t *testing.T) {
		_, err := setlocale("C")
		assert.NoError(t, err, "should be able to set locale to C")
		result := localeconv()
		assert.NotEmpty(t, result, "resulting lconv should not be empty")
		fmt.Println(result)
	})
}
