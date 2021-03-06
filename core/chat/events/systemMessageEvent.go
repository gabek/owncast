package events

import "github.com/owncast/owncast/core/data"

// SystemMessageEvent is a message displayed in chat on behalf of the server.
type SystemMessageEvent struct {
	Event
	MessageEvent
}

// SystemMessageEvent will return the object to send to all chat users.
func (e *SystemMessageEvent) GetBroadcastPayload() EventPayload {
	return EventPayload{
		"id":        e.Id,
		"timestamp": e.Timestamp,
		"body":      e.Body,
		"type":      SystemMessageSent,
		"user": EventPayload{
			"displayName": data.GetServerName(),
		},
	}
}

func (e *SystemMessageEvent) GetMessageType() EventType {
	return SystemMessageSent
}
