//go:build !wasm

package locale

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

var (
	localePattern = regexp.MustCompile(`^(?P<language>[[:lower:]]{2,3})(?:[-_])(?P<country>[[:alpha:]]{2})(?:.?)(?P<encoding>\S+)?$`)
	utf8Pattern   = regexp.MustCompile(`(?<utf8>[uU][tT][fF].?[8])`)
)

var (
	installedLocales = []string{}
	// localeMapping contains the short locale code like en_US to the code that
	// the system is using, which may be en_US.UTF-8 or something else.
	localeMapping = map[string]string{}
	// currencyMapping is a map keyed by the international currency code with the
	// value being an array of locale's that use that currency.
	currencyMapping = map[string][]string{}
)

func init() {
	locales := listLocales()
	for i := range locales {
		locale := locales[i]

		matches := localePattern.FindStringSubmatch(locale)
		grouped := make(map[string]string, 3)
		for x, group := range localePattern.SubexpNames() {
			if x != 0 && group != "" && len(matches) > x {
				grouped[group] = matches[x]
			}
		}

		// At the moment we only want to support UTF-8 locales.
		if !utf8Pattern.MatchString(grouped["encoding"]) {
			continue
		}

		// If we cannot determine the language or the country then skip the locale.
		if _, ok := grouped["language"]; !ok {
			continue
		}

		if _, ok := grouped["country"]; !ok {
			continue
		}

		// Make sure the locale is installed.
		if _, err := setlocale(locale); err != nil {
			// Locale is not installed!
			continue
		}

		shortCode := fmt.Sprintf("%s_%s", grouped["language"], grouped["country"])
		installedLocales = append(installedLocales, shortCode)
		localeMapping[shortCode] = locale

		currency, err := GetLConv(shortCode)
		if err != nil {
			continue
		}

		currencyCode := string(bytes.TrimSpace(currency.IntCurrSymbol))
		if currencyLocales, ok := currencyMapping[currencyCode]; ok {
			currencyMapping[currencyCode] = append(currencyLocales, shortCode)
		} else {
			currencyMapping[currencyCode] = []string{shortCode}
		}
	}
	sort.Strings(installedLocales)
}

// GetInstalledLocales will return an array of locales that are accepted by the
// other locale functions in this package for the current host system. The
// locale names returned here are standardized into a format such as `en_US` and
// will not include the unicode suffix.
func GetInstalledLocales() []string {
	return installedLocales
}

func adjustLocale(input string) string {
	matches := localePattern.FindStringSubmatch(input)
	grouped := make(map[string]string, 3)
	for i, group := range localePattern.SubexpNames() {
		if i != 0 && group != "" && len(matches) > i {
			grouped[group] = matches[i]
		}
	}

	localeCleaned := fmt.Sprintf("%s_%s", grouped["language"], grouped["country"])
	// The locale cleaned does not contain an encoding, so check our mapping to
	// see what the locale code should be with an encoding suffix. If we have one
	// use that otherwise use the provided locale code.
	if original, ok := localeMapping[localeCleaned]; ok {
		return original
	}

	return localeCleaned
}

// listLocalesCommand is a way to list installed locales on all operating
// systems. However on Windows, if this command is not available then we will
// fallback to doing on OS function call to list locales.
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
