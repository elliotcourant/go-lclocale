package locale

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetLocale(t *testing.T) {
	err := setLocale("bogus")
	assert.EqualError(t, err, "failed to set locale to: bogus")

	err = setLocale("C")
	assert.NoError(t, err, "should be able to set locale to C")
	assert.Equal(t, "C", getLocale())
}

func TestLocaleConv(t *testing.T) {
	err := setLocale("C")
	assert.NoError(t, err, "should be able to set locale to C")
	result := localeconv()
	assert.NotEmpty(t, result, "resulting lconv should not be empty")
	fmt.Println(result)
}
