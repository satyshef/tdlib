// AUTOGENERATED - DO NOT EDIT

package tdlib

// Stickers Represents a list of stickers
type Stickers struct {
	tdCommon
	Stickers []Sticker `json:"stickers"` // List of stickers
}

// MessageType return the string telegram-type of Stickers
func (stickers *Stickers) MessageType() string {
	return "stickers"
}

// NewStickers creates a new Stickers
//
// @param stickers List of stickers
func NewStickers(stickers []Sticker) *Stickers {
	stickersTemp := Stickers{
		tdCommon: tdCommon{Type: "stickers"},
		Stickers: stickers,
	}

	return &stickersTemp
}
