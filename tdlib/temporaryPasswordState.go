// AUTOGENERATED - DO NOT EDIT

package tdlib

// TemporaryPasswordState Returns information about the availability of a temporary password, which can be used for payments
type TemporaryPasswordState struct {
	tdCommon
	HasPassword bool  `json:"has_password"` // True, if a temporary password is available
	ValidFor    int32 `json:"valid_for"`    // Time left before the temporary password expires, in seconds
}

// MessageType return the string telegram-type of TemporaryPasswordState
func (temporaryPasswordState *TemporaryPasswordState) MessageType() string {
	return "temporaryPasswordState"
}

// NewTemporaryPasswordState creates a new TemporaryPasswordState
//
// @param hasPassword True, if a temporary password is available
// @param validFor Time left before the temporary password expires, in seconds
func NewTemporaryPasswordState(hasPassword bool, validFor int32) *TemporaryPasswordState {
	temporaryPasswordStateTemp := TemporaryPasswordState{
		tdCommon:    tdCommon{Type: "temporaryPasswordState"},
		HasPassword: hasPassword,
		ValidFor:    validFor,
	}

	return &temporaryPasswordStateTemp
}
