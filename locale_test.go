package locale

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetLocale(t *testing.T) {
	setLocale("en_US")
	assert.Equal(t, "en_US", getLocale())
	setLocale("da_DK")
	assert.Equal(t, "da_DK", getLocale())
}

func TestLocaleConv(t *testing.T) {
	setLocale("en_US")
	result := localeconv()
	fmt.Println(result)
}
