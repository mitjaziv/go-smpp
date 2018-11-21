// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
package pdutext

import (
	"golang.org/x/text/transform"

	"github.com/fiorix/go-smpp/smpp/encoding"
)

// GSM 7-bit (Spanish)
type GSM7Spanish []byte

// Type implements the Codec interface.
func (s GSM7Spanish) Type() DataCoding {
	return DefaultType
}

// Encode to GSM 7-bit (Spanish)
func (s GSM7Spanish) Encode() []byte {
	e := encoding.GSM7(
		false,
		encoding.Language(encoding.Spanish),
	).NewEncoder()

	es, _, err := transform.Bytes(e, s)
	if err != nil {
		return s
	}
	return es
}

// Decode from GSM 7-bit (Spanish)
func (s GSM7Spanish) Decode() []byte {
	e := encoding.GSM7(
		false,
		encoding.Language(encoding.Spanish),
	).NewDecoder()

	es, _, err := transform.Bytes(e, s)
	if err != nil {
		return s
	}
	return es
}
