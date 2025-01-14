//go:build !wasm

package locale_test

import (
	"testing"

	locale "github.com/elliotcourant/go-lclocale"
	"github.com/stretchr/testify/assert"
)

func TestGetInstalledLocales(t *testing.T) {
	t.Run("en_US must be installed", func(t *testing.T) {
		result := locale.GetInstalledLocales()
		assert.Contains(t, result, "en_US", "must contain the en_US locale")
	})
}
