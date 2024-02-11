package locale

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetLocale(t *testing.T) {
	setLocale("en_US.utf8")
	assert.Equal(t, "en_US.utf8", getLocale())
	setLocale("C")
	assert.Equal(t, "C", getLocale())
}

func TestLocaleConv(t *testing.T) {
	setLocale("C")
	result := localeconv()
	fmt.Println(result)
}
