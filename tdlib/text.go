// AUTOGENERATED - DO NOT EDIT

package tdlib

// Text Contains some text
type Text struct {
	tdCommon
	Text string `json:"text"` // Text
}

// MessageType return the string telegram-type of Text
func (text *Text) MessageType() string {
	return "text"
}

// NewText creates a new Text
//
// @param text Text
func NewText(text string) *Text {
	textTemp := Text{
		tdCommon: tdCommon{Type: "text"},
		Text:     text,
	}

	return &textTemp
}