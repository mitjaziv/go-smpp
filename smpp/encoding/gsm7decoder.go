package encoding

import (
	"bytes"
	"errors"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

var (
	// ErrInvalidByte means that a given byte is outside of the GSM 7-bit encoding range,
	// this can only happen during decoding.
	ErrInvalidByte = errors.New("invalid gsm7 byte")
)

type (
	gsm7Decoder struct {
		packed bool
		lang   Lang
	}
)

func (g gsm7Encoding) NewDecoder() *encoding.Decoder {
	return &encoding.Decoder{
		Transformer: &gsm7Decoder{
			packed: g.packed,
			lang:   g.lang,
		},
	}
}

func (g *gsm7Decoder) Reset() {
	/* not needed */
}

func (g *gsm7Decoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	if len(src) == 0 {
		return 0, 0, nil
	}

	septets := src
	if g.packed {
		septets = make([]byte, 0, len(src))
		count := 0
		remain := len(src) - count
		for remain > 0 {
			// Unpack by converting octets into septets.
			if remain >= 7 {
				septets = append(septets, src[count+0]&0x7F<<0)
				septets = append(septets, (src[count+1]&0x3F<<1)|(src[count+0]&0x80>>7))
				septets = append(septets, (src[count+2]&0x1F<<2)|(src[count+1]&0xC0>>6))
				septets = append(septets, (src[count+3]&0x0F<<3)|(src[count+2]&0xE0>>5))
				septets = append(septets, (src[count+4]&0x07<<4)|(src[count+3]&0xF0>>4))
				septets = append(septets, (src[count+5]&0x03<<5)|(src[count+4]&0xF8>>3))
				septets = append(septets, (src[count+6]&0x01<<6)|(src[count+5]&0xFC>>2))
				if src[count+6] > 0 {
					septets = append(septets, src[count+6]&0xFE>>1)
				}
				count += 7
			} else if remain >= 6 {
				septets = append(septets, src[count+0]&0x7F<<0)
				septets = append(septets, (src[count+1]&0x3F<<1)|(src[count+0]&0x80>>7))
				septets = append(septets, (src[count+2]&0x1F<<2)|(src[count+1]&0xC0>>6))
				septets = append(septets, (src[count+3]&0x0F<<3)|(src[count+2]&0xE0>>5))
				septets = append(septets, (src[count+4]&0x07<<4)|(src[count+3]&0xF0>>4))
				septets = append(septets, (src[count+5]&0x03<<5)|(src[count+4]&0xF8>>3))
				count += 6
			} else if remain >= 5 {
				septets = append(septets, src[count+0]&0x7F<<0)
				septets = append(septets, (src[count+1]&0x3F<<1)|(src[count+0]&0x80>>7))
				septets = append(septets, (src[count+2]&0x1F<<2)|(src[count+1]&0xC0>>6))
				septets = append(septets, (src[count+3]&0x0F<<3)|(src[count+2]&0xE0>>5))
				septets = append(septets, (src[count+4]&0x07<<4)|(src[count+3]&0xF0>>4))
				count += 5
			} else if remain >= 4 {
				septets = append(septets, src[count+0]&0x7F<<0)
				septets = append(septets, (src[count+1]&0x3F<<1)|(src[count+0]&0x80>>7))
				septets = append(septets, (src[count+2]&0x1F<<2)|(src[count+1]&0xC0>>6))
				septets = append(septets, (src[count+3]&0x0F<<3)|(src[count+2]&0xE0>>5))
				count += 4
			} else if remain >= 3 {
				septets = append(septets, src[count+0]&0x7F<<0)
				septets = append(septets, (src[count+1]&0x3F<<1)|(src[count+0]&0x80>>7))
				septets = append(septets, (src[count+2]&0x1F<<2)|(src[count+1]&0xC0>>6))
				count += 3
			} else if remain >= 2 {
				septets = append(septets, src[count+0]&0x7F<<0)
				septets = append(septets, (src[count+1]&0x3F<<1)|(src[count+0]&0x80>>7))
				count += 2
			} else if remain >= 1 {
				septets = append(septets, src[count+0]&0x7F<<0)
				count += 1
			} else {
				break
			}
			remain = len(src) - count
		}
	}

	// Get reverse lookup table and reverse escape table according to the language.
	reverseLookup, reverseEscape := getReverseSet(g.lang)

	nSeptet := 0
	builder := bytes.NewBufferString("")
	for nSeptet < len(septets) {
		b := septets[nSeptet]
		if b == escapeSequence {
			nSeptet++
			if nSeptet >= len(septets) {
				return 0, 0, ErrInvalidByte
			}
			e := septets[nSeptet]
			if r, ok := reverseEscape[e]; ok {
				builder.WriteRune(r)
			} else {
				return 0, 0, ErrInvalidByte
			}
		} else if r, ok := reverseLookup[b]; ok {
			builder.WriteRune(r)
		} else {
			return 0, 0, ErrInvalidByte
		}
		nSeptet++
	}
	text := builder.Bytes()
	nDst = len(text)

	if len(dst) < nDst {
		return 0, 0, transform.ErrShortDst
	}

	for x, b := range text {
		dst[x] = b
	}
	return nDst, nSrc, err
}
