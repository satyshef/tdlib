// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
	"fmt"
)

// LoginURLInfo Contains information about an inline button of type inlineKeyboardButtonTypeLoginUrl
type LoginURLInfo interface {
	GetLoginURLInfoEnum() LoginURLInfoEnum
}

// LoginURLInfoEnum Alias for abstract LoginURLInfo 'Sub-Classes', used as constant-enum here
type LoginURLInfoEnum string

// LoginURLInfo enums
const ()

func unmarshalLoginURLInfo(rawMsg *json.RawMessage) (LoginURLInfo, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch LoginURLInfoEnum(objMap["@type"].(string)) {

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// LoginURLInfoOpen An HTTP url needs to be open
type LoginURLInfoOpen struct {
	tdCommon
	URL         string `json:"url"`          // The URL to open
	SkipConfirm bool   `json:"skip_confirm"` // True, if there is no need to show an ordinary open URL confirm
}

// MessageType return the string telegram-type of LoginURLInfoOpen
func (loginURLInfoOpen *LoginURLInfoOpen) MessageType() string {
	return "loginUrlInfoOpen"
}

// NewLoginURLInfoOpen creates a new LoginURLInfoOpen
//
// @param uRL The URL to open
// @param skipConfirm True, if there is no need to show an ordinary open URL confirm
func NewLoginURLInfoOpen(uRL string, skipConfirm bool) *LoginURLInfoOpen {
	loginURLInfoOpenTemp := LoginURLInfoOpen{
		tdCommon:    tdCommon{Type: "loginUrlInfoOpen"},
		URL:         uRL,
		SkipConfirm: skipConfirm,
	}

	return &loginURLInfoOpenTemp
}

// LoginURLInfoRequestConfirmation An authorization confirmation dialog needs to be shown to the user
type LoginURLInfoRequestConfirmation struct {
	tdCommon
	URL                string `json:"url"`                  // An HTTP URL to be opened
	Domain             string `json:"domain"`               // A domain of the URL
	BotUserID          int64  `json:"bot_user_id"`          // User identifier of a bot linked with the website
	RequestWriteAccess bool   `json:"request_write_access"` // True, if the user needs to be requested to give the permission to the bot to send them messages
}

// MessageType return the string telegram-type of LoginURLInfoRequestConfirmation
func (loginURLInfoRequestConfirmation *LoginURLInfoRequestConfirmation) MessageType() string {
	return "loginUrlInfoRequestConfirmation"
}

// NewLoginURLInfoRequestConfirmation creates a new LoginURLInfoRequestConfirmation
//
// @param uRL An HTTP URL to be opened
// @param domain A domain of the URL
// @param botUserID User identifier of a bot linked with the website
// @param requestWriteAccess True, if the user needs to be requested to give the permission to the bot to send them messages
func NewLoginURLInfoRequestConfirmation(uRL string, domain string, botUserID int64, requestWriteAccess bool) *LoginURLInfoRequestConfirmation {
	loginURLInfoRequestConfirmationTemp := LoginURLInfoRequestConfirmation{
		tdCommon:           tdCommon{Type: "loginUrlInfoRequestConfirmation"},
		URL:                uRL,
		Domain:             domain,
		BotUserID:          botUserID,
		RequestWriteAccess: requestWriteAccess,
	}

	return &loginURLInfoRequestConfirmationTemp
}
