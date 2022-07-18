// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/satyshef/go-tdlib/tdlib"
)

// GetCommands Returns the list of commands supported by the bot for the given user scope and language; for bots only
// @param scope The scope to which the commands are relevant
// @param languageCode A two-letter ISO 639-1 country code or an empty string
func (client *Client) GetCommands(scope tdlib.BotCommandScope, languageCode string) (*tdlib.BotCommands, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":         "getCommands",
		"scope":         scope,
		"language_code": languageCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var botCommands tdlib.BotCommands
	err = json.Unmarshal(result.Raw, &botCommands)
	return &botCommands, err

}