// AUTOGENERATED - DO NOT EDIT

package tdlib

// PassportElements Contains information about saved Telegram Passport elements
type PassportElements struct {
	tdCommon
	Elements []PassportElement `json:"elements"` // Telegram Passport elements
}

// MessageType return the string telegram-type of PassportElements
func (passportElements *PassportElements) MessageType() string {
	return "passportElements"
}

// NewPassportElements creates a new PassportElements
//
// @param elements Telegram Passport elements
func NewPassportElements(elements []PassportElement) *PassportElements {
	passportElementsTemp := PassportElements{
		tdCommon: tdCommon{Type: "passportElements"},
		Elements: elements,
	}

	return &passportElementsTemp
}
