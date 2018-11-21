package encoding

import (
	"golang.org/x/text/encoding"
)

const escapeSequence = 0x1B

type (
	gsm7Encoding struct {
		packed bool
		lang   Lang
	}

	Option func(*gsm7Encoding)
)

// GSM7 returns a GSM 7-bit Bit Encoding.
//
// Set the packed flag to true if you wish to convert septets to octets,
// this should be false for most SMPP providers.
func GSM7(packed bool, opts ...Option) encoding.Encoding {
	enc := &gsm7Encoding{
		packed: packed,
	}

	// Apply optional parameters.
	applyDefaults(enc)
	for _, o := range opts {
		o(enc)
	}

	return enc
}

// applyDefaults values,
func applyDefaults(e *gsm7Encoding) {
	e.lang = Basic
}

// Language configures basic and extended charset tables.
func Language(lang Lang) Option {
	return func(e *gsm7Encoding) {
		e.lang = lang
	}
}

func (g gsm7Encoding) String() string {
	if g.packed {
		return "GSM 7-bit (Packed)"
	}
	return "GSM 7-bit (Unpacked)"
}

// Returns the characters, in the given text, that can not be represented in GSM 7-bit encoding.
func ValidateGSM7String(text string) []rune {
	forwardLookup := basicForward
	forwardEscape := spanishExtensionForward

	invalidChars := make([]rune, 0, 4)
	for _, r := range text {
		if _, ok := forwardLookup[r]; !ok {
			if _, ok := forwardEscape[r]; !ok {
				invalidChars = append(invalidChars, r)
			}
		}
	}
	return invalidChars
}

// Returns the bytes, in the given buffer, that are outside of the GSM 7-bit encoding range.
func ValidateGSM7Buffer(buffer []byte) []byte {
	reverseLookup := basicReverse
	reverseEscape := spanishExtensionReverse

	invalidBytes := make([]byte, 0, 4)
	count := 0
	for count < len(buffer) {
		b := buffer[count]
		if b == escapeSequence {
			count++
			if count >= len(buffer) {
				invalidBytes = append(invalidBytes, b)
				break
			}
			e := buffer[count]
			if _, ok := reverseEscape[e]; !ok {
				invalidBytes = append(invalidBytes, b, e)
			}
		} else if _, ok := reverseLookup[b]; !ok {
			invalidBytes = append(invalidBytes, b)
		}
		count++
	}
	return invalidBytes
}
