//go:build windows

package locale

import (
	"strings"
	"syscall"
	"unsafe"
)

var (
	// Load the required DLL
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	// Get a handle to the EnumSystemLocalesEx function
	procEnumSystemLocalesEx = kernel32.NewProc("EnumSystemLocalesEx")
)

func listLocales() []string {
	locales := make([]string, 0)
	enumLocalesCallback := func(
		locale *uint16,
		flags uint32,
		lParam uintptr,
	) uintptr {
		// Convert the UTF-16 string provided by the API to a Go string
		localeStr := syscall.UTF16ToString((*[1 << 29]uint16)(unsafe.Pointer(locale))[:])

		// If the locale is something like `ar` then skip it, we need the two or
		// three part locales like en-US and others
		if !strings.Contains(localeStr, "-") {
			return 1
		}

		locales = append(locales, strings.ReplaceAll(localeStr, "-", "_"))
		return 1 // Returning 1 continues enumeration
	}

	// Call EnumSystemLocalesEx with our callback function
	ret, _, _ := procEnumSystemLocalesEx.Call(
		syscall.NewCallback(enumLocalesCallback),
		0, // LOCALE_ALL, use appropriate flag as per your requirement
		0, // lParam, pass additional data to callback if needed
	)
	if ret == 0 {
		panic("EnumSystemLocalesEx failed!")
	}

	return locales
}
