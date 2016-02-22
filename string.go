package util

import "strconv"
import "strings"

func Str(nfNumeric float64, nLenght int, nDecimal int) string {

	return strconv.FormatFloat(nfNumeric, 'f', nDecimal, 64)[0:nLenght]

}

func Substr(cString string, nStart, nLenght int) string {
	return cString[nStart-1 : nLenght]
}

func Strextract(cString, cADelim, cCDelim string, nOccurrence int) string {

	return strings.Split(strings.Split(cString, cADelim)[nOccurrence], cCDelim)[0]
}
