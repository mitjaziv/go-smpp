// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
package pdutext

import (
	"bytes"
	"testing"
)

func TestGSM7SpanishEncoder(t *testing.T) {
	want := []byte(
		"\x48\x65\x6c\x6c\x6f\x20\x77\x6f\x72\x6c\x64\x20\x1b\x09\x20\x1b\x41\x20\x1b\x49\x20" +
			"\x1b\x4f\x20\x1b\x55\x20\x1b\x61\x20\x1b\x65\x20\x1b\x69\x20\x1b\x6f\x20\x1b\x75",
	)
	text := []byte("Hello world ç Á Í Ó Ú á € í ó ú")

	s := GSM7Spanish(text)
	if s.Type() != 0x00 {
		t.Fatalf("Unexpected data type; want 0x00, have %d", s.Type())
	}

	have := s.Encode()
	if !bytes.Equal(want, have) {
		t.Fatalf("Unexpected text; want %q, have %q", want, have)
	}
}

func TestGSM7SpanishDecoder(t *testing.T) {
	want := []byte("Hello world ç Á Í Ó Ú á € í ó ú")
	text := []byte(
		"\x48\x65\x6c\x6c\x6f\x20\x77\x6f\x72\x6c\x64\x20\x1b\x09\x20\x1b\x41\x20\x1b\x49\x20" +
			"\x1b\x4f\x20\x1b\x55\x20\x1b\x61\x20\x1b\x65\x20\x1b\x69\x20\x1b\x6f\x20\x1b\x75",
	)

	s := GSM7Spanish(text)
	if s.Type() != 0x00 {
		t.Fatalf("Unexpected data type; want 0x00, have %d", s.Type())
	}

	have := s.Decode()
	if !bytes.Equal(want, have) {
		t.Fatalf("Unexpected text; want %q, have %q", want, have)
	}
}
