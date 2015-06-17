package set1

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

//Challenge 1
//Convert hex to base64
func HtoB64(hexstr string) (string, error) {
	hb, err := hex.DecodeString(hexstr)
	if err != nil {
		return "", err
	}
	str:=base64.StdEncoding.EncodeToString(hb)
	return str, nil
}

//Challenge 2
//XOR Enc
func XOREnc(hex1 string, hex2 string) (string, error) {
	hb1, err1 := hex.DecodeString(hex1)
	hb2, err2 := hex.DecodeString(hex2)
	if (len(hb1) != len(hb2)) {
		return "", errors.New("String lengths incorrect length")
	}
	if err1 != nil {
		return "", err1
	}
	if err2 != nil {
		return "", err2
	}
	for i, _ := range hb1 {
		hb2[i] ^= hb1[i]
	}
	return hex.EncodeToString(hb2), nil
}

//Challenge 3
//XOR attack with char
func XORChar(ct string, key byte) (string, error) {
	ctb, err := hex.DecodeString(ct)
	if err != nil {
		return "", err
	}
	for i := range ctb {
		ctb[i] ^= key
	}
	return string(ctb), nil
}
