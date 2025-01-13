//go:build !wasm && darwin

package locale

var (
	localeLookupPaths = []string{
		"/usr/share/locale",
	}
	encodingPriority = []string{
		"utf-8",
		"iso8859-15",
		"iso8859-1",
		"", // No encoding specified is lowest priority
	}
)
