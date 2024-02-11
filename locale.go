package locale

/*
#include <stdlib.h>
#include <locale.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func setLocale(locale string) error {
	cLocale := C.CString(locale)
	defer C.free(unsafe.Pointer(cLocale))
	res := C.GoString(C.setlocale(C.LC_ALL, cLocale))
	if res == "" {
		return fmt.Errorf("failed to set locale to: %s", locale)
	}
	return nil
}

func getLocale() string {
	ptr := C.setlocale(C.LC_ALL, nil)
	currentLocale := C.GoString(ptr)
	return currentLocale
}

type LConv struct {
	DecimalPoint    []byte
	ThousandsSep    []byte
	Grouping        []byte
	IntCurrSymbol   []byte
	CurrencySymbol  []byte
	MonDecimalPoint []byte
	MonThousandsSep []byte
	MonGrouping     []byte
	PositiveSign    []byte
	NegativeSign    []byte
	PCSPrecedes     bool
	PSepBySpace     bool
	NCSPrecedes     bool
	NSepBySpace     bool
	PSignPosn       uint8
	NSignPosn       uint8
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
		PCSPrecedes:     byte(C.char(clconv.p_cs_precedes)) == byte(1),
		PSepBySpace:     byte(C.char(clconv.p_sep_by_space)) == byte(1),
		NCSPrecedes:     byte(C.char(clconv.n_cs_precedes)) == byte(1),
		NSepBySpace:     byte(C.char(clconv.n_sep_by_space)) == byte(1),
		PSignPosn:       uint8(C.char(clconv.p_sign_posn)),
		NSignPosn:       uint8(C.char(clconv.n_sign_posn)),
	}

	return lconv
}
