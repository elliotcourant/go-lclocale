//go:build !wasm

package locale_test

import (
	"sync"
	"testing"

	locale "github.com/elliotcourant/go-lclocale"
	"github.com/stretchr/testify/assert"
)

func TestGetLConv(t *testing.T) {
	t.Run("en_US", func(t *testing.T) {
		lconv, err := locale.GetLConv("en_US")
		assert.NoError(t, err, "should not return an error for en_US")
		assert.Subset(t, lconv.IntCurrSymbol, []byte("USD"))
		assert.Equal(t, []byte("."), lconv.DecimalPoint)
		assert.Equal(t, []byte(","), lconv.ThousandsSep)
		assert.Equal(t, []byte("$"), lconv.CurrencySymbol)
		assert.EqualValues(t, 2, lconv.FracDigits)
	})

	t.Run("ja_JP", func(t *testing.T) {
		lconv, err := locale.GetLConv("ja_JP")
		assert.NoError(t, err, "should not return an error for ja_JP")
		assert.Subset(t, lconv.IntCurrSymbol, []byte("JPY"))
		assert.Equal(t, []byte("."), lconv.DecimalPoint)
		assert.Equal(t, []byte(","), lconv.ThousandsSep)
		assert.EqualValues(t, 0, lconv.FracDigits)
	})

	t.Run("thread safety", func(t *testing.T) {
		// This test should be run with the --race flag. It will make sure that we
		// do not run into a race condition when we are calling the same locale from
		// multiple threads regarding the cache layer.
		numberOfThreads := 8
		wg := sync.WaitGroup{}
		wg.Add(numberOfThreads)

		startWg := sync.WaitGroup{}
		startWg.Add(1)

		for i := 0; i < numberOfThreads; i++ {
			go func() {
				// Wait until we are ready to start. This way they all run at once.
				startWg.Wait()
				defer wg.Done()
				lconv, err := locale.GetLConv("en_US")
				assert.NoError(t, err, "should not return an error for en_US")
				assert.Equal(t, []byte("."), lconv.DecimalPoint)
				assert.Equal(t, []byte(","), lconv.ThousandsSep)
				assert.Equal(t, []byte("$"), lconv.CurrencySymbol)
				assert.EqualValues(t, 2, lconv.FracDigits)
			}()
		}

		// Trigger all threads to start work
		startWg.Done()

		// Wait for all threads to finish
		wg.Wait()
	})
}
