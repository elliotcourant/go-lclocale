package locale

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetLocale(t *testing.T) {
	setLocale("POSIX")
	assert.Equal(t, "POSIX", getLocale())
	setLocale("C")
	assert.Equal(t, "C", getLocale())
}

func TestLocaleConv(t *testing.T) {
	setLocale("C")
	result := localeconv()
	fmt.Println(result)
}
