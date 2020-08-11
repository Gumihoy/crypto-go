package base64

var RFC4648 = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
var RFC4648_URLSAFE = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

type Padding rune

const (
	STANDARD Padding = '='
	NO       Padding = -1
)

// https://en.wanweibaike.com/wiki-BASE64
type Encoding struct {
	alphabet []byte
	pad      Padding
}

func New() *Encoding {
	return &Encoding{alphabet: RFC4648, pad: STANDARD}
}

func NewURL() *Encoding {
	return &Encoding{alphabet: RFC4648_URLSAFE, pad: STANDARD}
}

func (enc *Encoding) Encode(src []byte) []byte {
	sl := len(src)
	if sl == 0 {
		return make([]byte, 0)
	}

	dl := enc.encodedLen(sl)
	dst := make([]byte, dl)
	enc.encode0(src, 0, sl, dst)
	return dst
}

func (enc *Encoding) EncodeToString(src []byte) string {
	dst := enc.Encode(src)
	return string(dst)
}

func (enc *Encoding) encode0(src []byte, off int, end int, dst []byte) int {
	sp, dp := off, 0
	sl := off + (end-off)/3*3
	for sp < sl {
		bits := uint(src[sp])<<16 | uint(src[sp+1])<<8 | uint(src[sp+2])

		dst[dp] = enc.alphabet[bits>>18&0x3F]
		dst[dp+1] = enc.alphabet[bits>>12&0x3F]
		dst[dp+2] = enc.alphabet[bits>>6&0x3F]
		dst[dp+3] = enc.alphabet[bits&0x3F]
		sp += 3
		dp += 4
	}

	if sp == end {
		return dp
	}

	bits := uint(src[sp]) << 16

	remain := end - sp
	if remain == 2 {
		sp += 1
		bits |= uint(src[sp]) << 8
	}

	dst[dp] = enc.alphabet[bits>>18&0x3F]
	dst[dp+1] = enc.alphabet[bits>>12&0x3F]

	switch remain {
	case 1:
		if enc.pad != NO {
			dst[dp+2] = byte(enc.pad)
			dst[dp+3] = byte(enc.pad)
		}
	case 2:
		dst[dp+2] = enc.alphabet[bits>>6&0x3F]
		if enc.pad != NO {
			dst[dp+3] = byte(enc.pad)
		}
	}

	dp += 4
	return dp
}

func (enc *Encoding) encodedLen(len int) int {
	if enc.pad == NO {
		return (len*8 + 5) / 6
	}
	return ((len + 2) / 3) * 4
}

func (enc *Encoding) Decode(src []byte) []byte {
	return nil
}

func (enc *Encoding) decode0(src []byte, off int, end int, dst []byte) int {
	sp, dp := off, 0
	sl := off + (end-off)/4*4
	for sp < sl {

	}
	return dp
}

func (enc *Encoding) DecodedLen(len int) int {
	if enc.pad == NO {
		return len * 6 / 8
	}
	return len * 4 / 3
}
