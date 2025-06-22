// Copyright 2025 Ganesh Anantwar (finitology)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
// type RawParams json.RawMessage
