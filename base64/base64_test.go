package base64

import (
	eBase64 "encoding/base64"
	"fmt"
	"testing"
)

func TestEncoding_Encode(t *testing.T) {
	s := "xx"
	enc := New()
	fmt.Println("s1:", enc.EncodeToString([]byte(s)))

	enc2 := eBase64.StdEncoding
	dst := make([]byte, enc2.EncodedLen(len(s)))
	enc2.Encode(dst, []byte(s))
	fmt.Println("s2:", string(dst))
}
