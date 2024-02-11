package locale

// SignPosition dictates where the negative and positive sign symbols should be
// placed when formatting a number or monetary amount as a string for a given
// locale.
type SignPosition uint8

const (
	// Parentheses should surround the quantity and currency symbol.
	SignPositionParentheses SignPosition = 0
	// The sign string should precede the quantity and currency symbol.
	SignPositionPreceedsQuantityAndCurrency SignPosition = 1
	// The sign string should succeed the quantity and currency symbol.
	SignPositionSucceedsQuantityAndCurrency SignPosition = 2
	// The sign string should immediately precede the currency symbol.
	SignPositionImmediatelyPreceedsCurrency SignPosition = 3
	// The sign string should immediately succeed the currency symbol.
	SignPositionImmediatelysucceedsCurrency SignPosition = 4
)

// LConv contains the numeric and monetary information for a given locale. This
// data can be used to parse or properly format numbers or monetary amounts to
// that locales specification.
type LConv struct {
	// Radix character.
	DecimalPoint []byte
	// Separator for digit groupps to left of radix character.
	ThousandsSep []byte
	// Each element is the number of digits in a group; elements with higher
	// indicies are further left. An element with value 255 means that no further
	// grouping is done. An element with value 0 means that the previous element
	// is used for all groups further left.
	Grouping []uint8
	// First three characters are a currency symbol from ISO4217. Fourth character
	// is the separator. Fifth character is '\0'.
	IntCurrSymbol []byte
	// Local currency symbol.
	CurrencySymbol []byte
	// Radix character.
	MonDecimalPoint []byte
	// Like ThousandsSep above.
	MonThousandsSep []byte
	// Like Grouping above.
	MonGrouping []uint8
	// Sign for positive values.
	PositiveSign []byte
	// Sign for negative values.
	NegativeSign []byte
	// International fractional digits.
	IntFracDigits uint8
	// Local fractional digits.
	FracDigits uint8
	// True if CurrencySymbol precedes a positive value, false if succeeds.
	PCSPrecedes bool
	// True if a space separates CurrencySymbol from a positive value.
	PSepBySpace bool
	// True if CurrencySymbol precedes a negative value, false if it succeeds.
	NCSPrecedes bool
	// True if a space separates the CurrencySymbol from a negative value.
	NSepBySpace bool
	PSignPosn   SignPosition
	NSignPosn   SignPosition
}

// GetLConv will query the current installed locales on the host system and
// return the lconv data for the specified locale if it is installed. If the
// locale is not installed then an error is returned. The locale name may be
// adjusted to make calling easier. For example; `en_US` may be corrected to
// `en_US.utf8` on Linux or `en_US.UTF-8` on Darwin depending on what base
// locales are installed on the system.
func GetLConv(locale string) (*LConv, error) {
	adjusted := adjustLocale(locale)

	// Because setLocale affects the entire process, we need to lock this whenever
	// we change the locale. This way we know that the local has been set to the
	// one we need and will not change while we are working with it.
	localeMutex.Lock()
	defer localeMutex.Unlock()
	if err := setLocale(adjusted); err != nil {
		return nil, err
	}

	result := localeconv()
	return &result, nil
}
