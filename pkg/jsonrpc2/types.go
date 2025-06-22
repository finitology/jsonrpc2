package jsonrpc2

import (
	"encoding/json"
	"fmt"
)

// ID can be string, number or null
type ID struct {
	Str *string
	Num *float64
}

// UnmarshalJSON implements json.Unmarshaler
func (id *ID) UnmarshalJSON(data []byte) error {
	// null ID
	if string(data) == "null" {
		return nil
	}

	// try string
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		id.Str = &s
		return nil
	}

	// try number
	var n float64
	if err := json.Unmarshal(data, &n); err == nil {
		id.Num = &n
		return nil
	}

	return fmt.Errorf("invalid ID type")
}

// MarshalJSON implements json.Marshaler
func (id ID) MarshalJSON() ([]byte, error) {
	if id.Str != nil {
		return json.Marshal(*id.Str)
	}
	if id.Num != nil {
		return json.Marshal(*id.Num)
	}
	return []byte("null"), nil
}

// Version always "2.0"
const Version = "2.0"

// RawParams is used for lazy decoding
type RawParams json.RawMessage
