//go:build darwin

package locale

import (
	"fmt"
	"testing"
)

func TestListLocalesDarwin(t *testing.T) {
	knownLocales := listLocales()
	for i := range knownLocales {
		knownLocal := knownLocales[i]
		if err := setLocale(knownLocal); err != nil {
			fmt.Println("Locale:", knownLocal, "is NOT installed")
		} else {
			fmt.Println("Locale:", knownLocal, "IS installed")
		}
	}
}
