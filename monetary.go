package locale

type SignPosition uint8

const (
	SignPositionParentheses                 SignPosition = 0
	SignPositionPreceedsQuantityAndCurrency SignPosition = 1
	SignPositionSucceedsQuantityAndCurrency SignPosition = 2
	SignPositionImmediatelyPreceedsCurrency SignPosition = 3
	SignPositionImmediatelysucceedsCurrency SignPosition = 4
)

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

func GetLConv(locale string) (*LConv, error) {
	localeMutex.Lock()
	defer localeMutex.Unlock()

	adjusted := adjustLocale(locale)

	if err := setLocale(adjusted); err != nil {
		return nil, err
	}

	result := localeconv()
	return &result, nil
}
