//go:build !wasm && unix && !darwin

package locale

var (
	localeLookupPaths = []string{
		"/usr/lib/locale",
	}
	encodingPriority = []string{
		"utf-8",
		"", // No encoding specified is lowest priority
	}
)
