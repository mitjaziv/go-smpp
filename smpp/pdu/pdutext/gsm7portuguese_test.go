// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
package pdutext

import (
	"bytes"
	"testing"
)

func TestGSM7PortugueseEncoder(t *testing.T) {
	want := []byte(
		"\x48\x65\x6c\x6c\x6f\x20\x77\x6f\x72\x6c\x64\x20\x1a\x20\x14\x20\x1c\x20\x1b\x12\x20\x1b\x13\x20\x16\x20" +
			"\x1b\x05\x20\x1b\x15\x20\x5d\x20\x18\x20\x1b\x75\x20\x1b\x16\x20\x1b\x17\x20\x1b\x18\x20\x09\x20\x1b" +
			"\x19\x20\x40\x20\x1b\x69\x20\x0b\x20\x1b\x5b\x20\x7b\x20\x0c\x20\x1b\x5c\x20\x7c\x20\x0e\x20\x0f\x20" +
			"\x1e\x20\x19\x20\x1b\x6f\x20\x1d",
	)
	text := []byte("Hello world | À Â Φ Γ ^ ê Ω Ú € ú Π Ψ Σ ç Θ Í í Ô Ã ã ô Õ õ Á á Ê Ó ó â")

	s := GSM7Portuguese(text)
	if s.Type() != 0x00 {
		t.Fatalf("Unexpected data type; want 0x00, have %d", s.Type())
	}

	have := s.Encode()
	if !bytes.Equal(want, have) {
		t.Fatalf("Unexpected text; want %q, have %q", want, have)
	}
}

func TestGSM7PortugueseDecoder(t *testing.T) {
	want := []byte("Hello world | À Â Φ Γ ^ ê Ω Ú € ú Π Ψ Σ ç Θ Í í Ô Ã ã ô Õ õ Á á Ê Ó ó â")
	text := []byte(
		"\x48\x65\x6c\x6c\x6f\x20\x77\x6f\x72\x6c\x64\x20\x1a\x20\x14\x20\x1c\x20\x1b\x12\x20\x1b\x13\x20\x16\x20" +
			"\x1b\x05\x20\x1b\x15\x20\x5d\x20\x18\x20\x1b\x75\x20\x1b\x16\x20\x1b\x17\x20\x1b\x18\x20\x09\x20\x1b" +
			"\x19\x20\x40\x20\x1b\x69\x20\x0b\x20\x1b\x5b\x20\x7b\x20\x0c\x20\x1b\x5c\x20\x7c\x20\x0e\x20\x0f\x20" +
			"\x1e\x20\x19\x20\x1b\x6f\x20\x1d",
	)

	s := GSM7Portuguese(text)
	if s.Type() != 0x00 {
		t.Fatalf("Unexpected data type; want 0x00, have %d", s.Type())
	}

	have := s.Decode()
	if !bytes.Equal(want, have) {
		t.Fatalf("Unexpected text; want %q, have %q", want, have)
	}
}
