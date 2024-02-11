//go:build windows

package locale

import (
	"fmt"
	"testing"
)

func TestListLocalesWindows(t *testing.T) {
	fmt.Println(listLocales())
}
