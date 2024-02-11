//go:build unix

package locale

func listLocales() []string {
	return listLocalesCommand()
}
