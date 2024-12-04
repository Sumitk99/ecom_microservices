// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/4fcf747dfafc951e1dcf3077327e3dcee9107db3

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// BulkError type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4fcf747dfafc951e1dcf3077327e3dcee9107db3/specification/security/_types/Bulk.ts#L24-L33
type BulkError struct {
	// Count The number of errors
	Count int `json:"count"`
	// Details Details about the errors, keyed by role name
	Details map[string]ErrorCause `json:"details"`
}

func (s *BulkError) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "details":
			if s.Details == nil {
				s.Details = make(map[string]ErrorCause, 0)
			}
			if err := dec.Decode(&s.Details); err != nil {
				return fmt.Errorf("%s | %w", "Details", err)
			}

		}
	}
	return nil
}

// NewBulkError returns a BulkError.
func NewBulkError() *BulkError {
	r := &BulkError{
		Details: make(map[string]ErrorCause, 0),
	}

	return r
}