// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
)

// MessageStatistics A detailed statistics about a message
type MessageStatistics struct {
	tdCommon
	MessageInteractionGraph StatisticalGraph `json:"message_interaction_graph"` // A graph containing number of message views and shares
}

// MessageType return the string telegram-type of MessageStatistics
func (messageStatistics *MessageStatistics) MessageType() string {
	return "messageStatistics"
}

// NewMessageStatistics creates a new MessageStatistics
//
// @param messageInteractionGraph A graph containing number of message views and shares
func NewMessageStatistics(messageInteractionGraph StatisticalGraph) *MessageStatistics {
	messageStatisticsTemp := MessageStatistics{
		tdCommon:                tdCommon{Type: "messageStatistics"},
		MessageInteractionGraph: messageInteractionGraph,
	}

	return &messageStatisticsTemp
}

// UnmarshalJSON unmarshal to json
func (messageStatistics *MessageStatistics) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	messageStatistics.tdCommon = tempObj.tdCommon

	fieldMessageInteractionGraph, _ := unmarshalStatisticalGraph(objMap["message_interaction_graph"])
	messageStatistics.MessageInteractionGraph = fieldMessageInteractionGraph

	return nil
}
