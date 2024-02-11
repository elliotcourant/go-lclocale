package locale

import (
	"os/exec"
	"sort"
	"strings"
)

func listLocalesCommand() []string {
	cmd := exec.Command("locale", "-a")
	output, err := cmd.Output()
	if err != nil {
		panic("'locale -a' failed!")
	}

	locales := strings.Split(string(output), "\n")
	dedupe := map[string]struct{}{}
	for i := range locales {
		locale := locales[i]
		locale = strings.TrimSpace(locale)
		if locale == "" {
			continue
		}
		parts := strings.SplitN(locale, "@", 2)

		// Trim things like an odd suffix on windows?
		if len(parts) > 1 {
			dedupe[parts[0]] = struct{}{}
		} else {
			dedupe[locale] = struct{}{}
		}
	}
	locales = make([]string, 0, len(dedupe))
	for locale := range dedupe {
		locales = append(locales, locale)
	}

	sort.Strings(locales)

	return locales
}
