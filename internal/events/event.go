package events

import "encoding/json"

type Type string

type Event struct {
	Type    Type            `json:"type"`
	Version int             `json:"version"`
	Data    json.RawMessage `json:"data"`
}
