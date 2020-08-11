package blowfish

type Blowfish struct {
	p [18]uint32
	s [4][256]uint32
}

func NewBlowfish(key []byte) *Blowfish {
	bf := new(Blowfish)
	keyLen := len(key)
	for i := 0; i < 18; i++ {
		bf.p[i] ^= uint32(key[i%keyLen])
	}

	var l, r uint32 = 0, 0

	for i := 0; i < 18; i += 2 {
		bf.encrypt(l, r)
		bf.p[i] = l
		bf.p[i+1] = r
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 256; j += 2 {
			bf.encrypt(l, r)
			bf.s[i][j] = l
			bf.s[i][j+1] = r
		}
	}
	return bf
}

// https://en.wanweibaike.com/wiki-Blowfish_(cipher)
func (b *Blowfish) encrypt(l uint32, r uint32) {
	for i := 0; i < 16; i += 2 {
		l ^= b.p[i]
		r ^= b.f(l)
		r ^= b.p[i+1]
		l ^= b.f(r)
	}
	l ^= b.p[16]
	r ^= b.p[17]
	l, r = r, l
}

func (b *Blowfish) f(x uint32) uint32 {
	h := b.s[0][x>>24] + b.s[1][x>>16&0xff]
	return (h ^ b.s[2][x>>8&0xff]) + b.s[3][x&0xff]
}

// https://en.wanweibaike.com/wiki-Blowfish_(cipher)
func (b *Blowfish) decrypt(l uint32, r uint32) {
	for i := 16; i > 0; i -= 2 {
		l ^= b.p[i+1]
		r ^= b.f(l)
		r ^= b.p[i]
		l ^= b.f(r)
	}
	l ^= b.p[1]
	r ^= b.p[0]

	l, r = r, l
}
