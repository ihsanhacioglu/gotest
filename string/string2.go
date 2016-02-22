package string

import (
	//	"bytes"
	//	"errors"
	"fmt"
	"strings"
)

const (
	// string.ascii_letters
	CLETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// string.ascii_lowercase
	CLOWERS = "abcdefghijklmnopqrstuvwxyz"

	// string.ascii_uppercase
	CUPPERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// string.digits
	CDIGITS = "0123456789"

	// string.hexdigits
	CHEXADE = "0123456789abcdefABCDEF"

	// string.punctuation
	CPUNCTUATION = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

	// string.printable
	CPPRINTABLE = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\x0b\x0c"

	// string.whitespace
	CWHITESPACE = " \t\n\r\x0b\x0c"
)

/*
string.IsDigit fonksiyonu cText sadece rakamlardan olusuyorsa true döndürür
TASK: Text içinde özel karakterler (space, v,rgül ve noktalama işareti varsa dikkate alınacak şekilde düzenlenmeli
*/

func IsDigit(cText string) bool {

	var lDigit = false // bool

	if Empty(cText) {
		return false
	}

	for ii, rHarf := range cText {

		fmt.Println(ii, rHarf, string(rHarf))

		if strings.IndexRune(CLETTERS, rHarf) > 0 {
			return false

		} else {
			if strings.IndexRune(CDIGITS, rHarf) > 0 {
				lDigit = true
			}
		}
	}

	return lDigit
}

func Empty(cText string) bool {
	return cText == ""
}

// CompressSpaces trims s and replaces consecutive spaces in s with a single space.

func CompressSpaces(cText string) string {
	return strings.Join(strings.Fields(cText), " ")
}
