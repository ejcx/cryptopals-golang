package set1_test
import (
	"bufio"
	"encoding/hex"
	"os"
	"regexp"
	"cryptopals-go/set1"
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
func looksenglish(s string) bool {
	re, _ := regexp.Compile("^[a-zA-Z'\\s!.\"?]+$")
	return re.MatchString(s)
}
func TestXORChar(t *testing.T) {
	ct := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	for i := 0 ; i < 255; i++ {
		r, _ :=	set1.XORChar(ct, byte(i))
		if looksenglish(r) {
			t.Log(r)
			return
		}
	}
	t.Fail()
}
//Challenge 4
//Brute force using the stuff from challenge 3
func TestFindXORdStr(t *testing.T){
	f, err := os.Open("4.txt")
	if err != nil {
		t.Fail()
	}
	br := bufio.NewReader(f)
	for line, p, err := br.ReadLine();; {
		if p || err != nil {
			break
		}
		for i:= 0 ; i < 255 ; i++ {
			s, _ :=	set1.XORChar(string(line), byte(i))
			if looksenglish(s){
				t.Logf("\nString: %sChar: %d", s, i)
				return
			}
		}
		line, p, err = br.ReadLine();
	}
	t.Fail()
}

func TestRepeatedXORCipher(t *testing.T){
	pt := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	win := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`
	key := "ICE"
	ct := set1.MultiTimePad([]byte(pt), []byte(key))
	res := hex.EncodeToString(ct)
	t.Logf("Result %s", res)
	if res != win {
		t.Fail()
	}
}
