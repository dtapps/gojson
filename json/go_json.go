// Copyright 2017 Bo-Yi Wu. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build go_json

package json

import json "github.com/goccy/go-json"

var (
	// Marshal is exported by gojson/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by gojson/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by gojson/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by gojson/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by gojson/json package.
	NewEncoder = json.NewEncoder
)