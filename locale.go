//go:build !wasm

package locale

/*
#include <stdlib.h>
#include <locale.h>
*/
import "C"
import (
	"fmt"
	"sync"
	"unsafe"
)

var (
	localeMutex = sync.Mutex{}
)

// Valid takes a locale code and will check to see if it is installed on the
// current system. If it is then it will return true. If the locale specified is
// either not valid or not installed then this will return false.
func Valid(locale string) bool {
	adjusted := adjustLocale(locale)

	// Because setLocale affects the entire process, we need to lock this whenever
	// we change the locale. This way we know that the local has been set to the
	// one we need and will not change while we are working with it.
	localeMutex.Lock()
	defer localeMutex.Unlock()
	if err := setLocale(adjusted); err != nil {
		return false
	}

	return true
}

func setLocale(locale string) error {
	cLocale := C.CString(locale)
	defer C.free(unsafe.Pointer(cLocale))
	res := C.GoString(C.setlocale(C.LC_ALL, cLocale))
	if res != locale {
		return fmt.Errorf("failed to set locale to: %s", locale)
	}
	return nil
}

func getLocale() string {
	ptr := C.setlocale(C.LC_ALL, nil)
	currentLocale := C.GoString(ptr)
	return currentLocale
}

func localeconv() LConv {
	clconv := C.localeconv()

	lconv := LConv{
		DecimalPoint:    []byte(C.GoString(clconv.decimal_point)),
		ThousandsSep:    []byte(C.GoString(clconv.thousands_sep)),
		Grouping:        []byte(C.GoString(clconv.grouping)),
		IntCurrSymbol:   []byte(C.GoString(clconv.int_curr_symbol)),
		CurrencySymbol:  []byte(C.GoString(clconv.currency_symbol)),
		MonDecimalPoint: []byte(C.GoString(clconv.mon_decimal_point)),
		MonThousandsSep: []byte(C.GoString(clconv.mon_thousands_sep)),
		MonGrouping:     []byte(C.GoString(clconv.mon_grouping)),
		PositiveSign:    []byte(C.GoString(clconv.positive_sign)),
		NegativeSign:    []byte(C.GoString(clconv.negative_sign)),
		IntFracDigits:   uint8(C.char(clconv.int_frac_digits)),
		FracDigits:      uint8(C.char(clconv.frac_digits)),
		PCSPrecedes:     byte(C.char(clconv.p_cs_precedes)) == byte(1),
		PSepBySpace:     byte(C.char(clconv.p_sep_by_space)) == byte(1),
		NCSPrecedes:     byte(C.char(clconv.n_cs_precedes)) == byte(1),
		NSepBySpace:     byte(C.char(clconv.n_sep_by_space)) == byte(1),
		PSignPosn:       SignPosition(C.char(clconv.p_sign_posn)),
		NSignPosn:       SignPosition(C.char(clconv.n_sign_posn)),
	}

	return lconv
}
