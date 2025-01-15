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

func setlocale(locale string) (string, error) {
	// Passing an empty locale will return the current locale the thread is using.
	if locale == "" {
		ptr := C.setlocale(C.LC_ALL, nil)
		currentLocale := C.GoString(ptr)
		return currentLocale, nil
	}

	// Otherwise we can send the specified locale directly.
	cLocale := C.CString(locale)
	defer C.free(unsafe.Pointer(cLocale))
	result := C.GoString(C.setlocale(C.LC_ALL, cLocale))
	// If the resulting locale code matches our input then we were able to
	// successfully switch locales, otherwise the specified locale is likely not
	// installed on this system.
	if result != locale {
		return "", fmt.Errorf("failed to set locale, wanted: [%s] got: [%s]", locale, result)
	}
	return result, nil
}
