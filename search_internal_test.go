//go:build !wasm

package locale

import (
	"fmt"
	"testing"
)

func TestListLocales(t *testing.T) {
	knownLocales := listLocales()
	for i := range knownLocales {
		knownLocal := knownLocales[i]
		if _, err := setlocale(knownLocal); err != nil {
			fmt.Println("Locale:", knownLocal, "is NOT installed")
		} else {
			fmt.Println("Locale:", knownLocal, "IS installed")
		}
	}
}
