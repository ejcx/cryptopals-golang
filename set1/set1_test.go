package set1_test
import (
	"regexp"
	"set1"
	"testing"
)
func TestHB64Test(t *testing.T) {
	r, err  := set1.HtoB64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if r != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" || err != nil {
		t.Fail()
	}
}
func TestXOREnc(t *testing.T) {
	s1 := "1c0111001f010100061a024b53535009181c"
	s2 := "686974207468652062756c6c277320657965"
	r, err  := set1.XOREnc(s1, s2)
	if r != "746865206b696420646f6e277420706c6179" || err != nil {
		t.Fail()
	}
}
func TestXORChar(t *testing.T) {
	ct := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	for i := 0 ; i < 255; i++ {
		r, _ :=	set1.XORChar(ct, byte(i))
		re, _ := regexp.Compile("^[a-zA-Z'\\s]+$")
		if re.MatchString(r) {
			t.Log(r)
			return
		}
	}
	t.Fail()
}
