// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
)

// ThemeSettings Describes theme settings
type ThemeSettings struct {
	tdCommon
	AccentColor        int32          `json:"accent_color"`         // Theme accent color in ARGB format
	Background         *Background    `json:"background"`           // The background to be used in chats; may be null
	MessageFill        BackgroundFill `json:"message_fill"`         // The fill to be used as a background for outgoing messages
	AnimateMessageFill bool           `json:"animate_message_fill"` // If true, the freeform gradient fill needs to be animated on every sent message
}

// MessageType return the string telegram-type of ThemeSettings
func (themeSettings *ThemeSettings) MessageType() string {
	return "themeSettings"
}

// NewThemeSettings creates a new ThemeSettings
//
// @param accentColor Theme accent color in ARGB format
// @param background The background to be used in chats; may be null
// @param messageFill The fill to be used as a background for outgoing messages
// @param animateMessageFill If true, the freeform gradient fill needs to be animated on every sent message
func NewThemeSettings(accentColor int32, background *Background, messageFill BackgroundFill, animateMessageFill bool) *ThemeSettings {
	themeSettingsTemp := ThemeSettings{
		tdCommon:           tdCommon{Type: "themeSettings"},
		AccentColor:        accentColor,
		Background:         background,
		MessageFill:        messageFill,
		AnimateMessageFill: animateMessageFill,
	}

	return &themeSettingsTemp
}

// UnmarshalJSON unmarshal to json
func (themeSettings *ThemeSettings) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		AccentColor        int32       `json:"accent_color"`         // Theme accent color in ARGB format
		Background         *Background `json:"background"`           // The background to be used in chats; may be null
		AnimateMessageFill bool        `json:"animate_message_fill"` // If true, the freeform gradient fill needs to be animated on every sent message
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	themeSettings.tdCommon = tempObj.tdCommon
	themeSettings.AccentColor = tempObj.AccentColor
	themeSettings.Background = tempObj.Background
	themeSettings.AnimateMessageFill = tempObj.AnimateMessageFill

	fieldMessageFill, _ := unmarshalBackgroundFill(objMap["message_fill"])
	themeSettings.MessageFill = fieldMessageFill

	return nil
}
