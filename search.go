package locale

import (
	"os/exec"
	"sort"
	"strings"
)

var (
	shortLocales     = []string{}
	installedLocales = []string{}
)

func init() {
	locales := listLocales()
	for i := range locales {
		locale := locales[i]
		if err := setLocale(locale); err != nil {
			// Locale is not installed!
			continue
		}

		installedLocales = append(installedLocales, locale)
	}
	dedupe := map[string]struct{}{}
	for i := range installedLocales {
		locale := installedLocales[i]
		parts := strings.SplitN(locale, ".", 2)
		dedupe[parts[0]] = struct{}{}
	}
	shortLocales = make([]string, 0, len(dedupe))
	for locale := range dedupe {
		shortLocales = append(shortLocales, locale)
	}
	sort.Strings(shortLocales)
}

// GetInstalledLocales will return an array of locales that are accepted by the
// other locale functions in this package for the current host system. The
// locale names returned here are standardized into a format such as `en_US` and
// will not include the unicode suffix.
func GetInstalledLocales() []string {
	return shortLocales
}

func adjustLocale(input string) string {
	result := input
	lowerInput := strings.ToLower(input)
	for i := range installedLocales {
		locale := installedLocales[i]
		if strings.EqualFold(locale, input) {
			return result
		}
		lowerLocale := strings.ToLower(locale)

		// If this locale has the input as a prefix then stage this locale to be
		// returned. This would be like if we had en_US as an input but en_US.UTF-8
		// is installed.
		if strings.HasPrefix(lowerLocale, lowerInput) {
			result = locale
		}
	}

	return result
}

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
