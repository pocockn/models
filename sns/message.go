package sns

import "encoding/json"

type (
	Message struct {
		ID      string           `json:"id"`
		Payload *json.RawMessage `json:"payload"`
	}
)
