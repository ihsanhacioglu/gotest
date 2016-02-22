package string

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func test() {

	s := "a11ba22ba333ba4444ba55555ba666666b"

	fmt.Println("StrExtract: ", StrExtract(s, "a", "b", 6))

	fmt.Println("Occur:", Occurs(s, "b"))
	fmt.Println("At:", At(s, "b"))
	fmt.Println("SubStr:", SubStr(s, 1, 2))
	fmt.Println("SubStrC:", SubStr(s, 1, 5))
	fmt.Println("Transform:", Transform(false))
	fmt.Println("ChrTran:", "1234", ChrTran("1214", "12", ""))
	fmt.Println("Asc:", Asc("本"))
	fmt.Println("Upper:", Upper("abcşğıiçüöæä本"))
	fmt.Println("Lower:", Lower(Upper("abcşğıiçüöæä本")))
	fmt.Println("Lowertr:", LowerTr(UpperTr("abcşğıiçüöæä本")))
}

/*
 Kısaltmalar :

a: array
r: rune
c: char
s: string
m: memo
o: object
o: struct
q: sql
t: table
d: date
z: Time
n: numeric
f: float
l: logical
x: Tipi Bilinmeyen
i: Interface
p: Pointer
?: map
f: fonksiyon



sMember := aAtext[1]   // string array üyesi
nMember := aAtext[1]   // integer array üyesi

dAtarih :=
dCtarih :=

*/

/* StrExtract Retrieves a string between two delimiters.

StrExtract(sExper, cAdelim, cCdelim, nOccur)

sExper:  Specifies the expression to search.
cAdelim: Specifies the character that delimits the beginning of sExper.
cCdelim: Specifies the character that delimits the end of sExper.
nOccur:  Specifies at which occurrence of cAdelim in sExper to start the extraction.

If cAdelim is an empty string, the search is conducted from the beginning of sExper to the first occurrence of cCdelim.
If cCdelim is an empty string StrExtract( ) returns a string from nOccurrence of cAdelim to the end of sExper.
*/

func StrExtract(sExper, sAdelim, sCdelim string, nOccur int) string {

	sExper = "." + sExper + "."
	sAdelim = Nvl(sAdelim, ".")
	sCdelim = StringNvl(sCdelim, ".")

	if nOccur == 0 {
		nOccur = 1
	}

	aExper := strings.Split(sExper, sAdelim)

	if len(aExper) <= nOccur {
		return ""
	}

	sMember := aExper[nOccur]
	aExper = strings.Split(sMember, sCdelim)

	if len(aExper) == 1 {
		return ""
	}

	return aExper[0]
}

func StringNvl(cAText, cRetVal string) string {

	if cAText == "" {
		cAText = cRetVal
	}

	return cAText
}

func Nvl(xAText, xRetVal interface{}) string {

	fmt.Println("Nvl:", xAText)

	xVal := reflect.ValueOf(xAText)
	typ := xVal.Type()
	Kind := xVal.Kind()

	fmt.Println("Val:", xVal, typ, Kind)

	if !xVal.IsValid() {
		return "<zero Value>"
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(xVal.Int(), 10)
	case reflect.String:
		return xVal.String()
	}

	/*cAText := string(iAText)

	if cAText == "" {
		cAText = cRetVal
	}*/

	return "" // cAText
}

/*
func Transform(zExprArgs interface{}) string {
	return toString(reflect.ValueOf(zExprArgs))
}

func toString(val reflect.Value) string {

	var str string
	if !val.IsValid() {
		return "<zero Value>"
	}
	typ := val.Type()
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'g', -1, 64)
	case reflect.Complex64, reflect.Complex128:
		c := val.Complex()
		return strconv.FormatFloat(real(c), 'g', -1, 64) + "+" + strconv.FormatFloat(imag(c), 'g', -1, 64) + "i"
	case reflect.String:
		return val.String()
	case reflect.Bool:
		if val.Bool() {
			return "true"
		} else {
			return "false"
		}
	case reflect.Ptr:
		v := val
		str = typ.String() + "("
		if v.IsNil() {
			str += "0"
		} else {
			str += "&" + toString(v.Elem())
		}
		str += ")"
		return str
	case reflect.Array, reflect.Slice:
		v := val
		str += typ.String()
		str += "{"
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				str += ", "
			}
			str += toString(v.Index(i))
		}
		str += "}"
		return str
	case reflect.Map:
		t := typ
		str = t.String()
		str += "{"
		str += "<can't iterate on maps>"
		str += "}"
		return str
	case reflect.Chan:
		str = typ.String()
		return str
	case reflect.Struct:
		m, ok := val.Type().MethodByName("String")
		if ok {
			str = m.Func.Call([]reflect.Value{val})[0].String()
		} else {

			t := typ
			v := val
			str += t.String()
			str += "{"
			for i, n := 0, v.NumField(); i < n; i++ {
				if i > 0 {
					str += ", "
				}
				str += toString(v.Field(i))
			}
			str += "}"
		}
		return str
	case reflect.Interface:
		return typ.String() + "(" + toString(val.Elem()) + ")"
	case reflect.Func:
		v := val
		return typ.String() + "(" + strconv.FormatUint(uint64(v.Pointer()), 10) + ")"
	default:
		panic("Transform: can't print type " + typ.String())
	}
	return "Transform: can't happen"
}

*/

//Returns the number of times a character expression occurs within another character expression.

func Occurs(cAText, cSubStr string) int {
	return strings.Count(cAText, cSubStr)
}

/*
Asc Returns the ANSI value for the leftmost character in a character expression.
if This value bigger than 255, Rune Value is obtained.

Asc(cExpression)

cExpression: Specifies the character expression containing the character whose ANSI value ASC( ) returns.
Any characters after the first character in cExpression are ignored by ASC( ).

Remarks: Asc( ) returns the position of the character in the character table of the current code page. Every character has a unique ANSI value in the range from 0 to 255.

Example: Asc("ABC")  -> 65

*/

func Asc(sAtext string) int { return int([]rune(sAtext)[0]) }

func At(cAText, cSubStr string) int { return strings.Index(cAText, cSubStr) + 1 }

/*
SubStr Returns a character string from the given character expression or memo field, starting at a specified position in the character expression or memo field and continuing for a specified number of characters.
SubStr(cAText, nStart [, nLenght])
cAText: Specifies the character expression or memo field from which the character string is returned.
nStart : Specifies the position in the character expression or memo field cExpression from where the character string is returned. The first character of cExpression is position 1.

*/

func SubStr(cAText string, nStart, nLenght int) string { return cAText[nStart-1 : nLenght] }

/*
SubStrC Returns a character string from the given character expression or memo field.

SubStrC(cExpression, nStartPosition [, nStrReturned])

cExpression:    Specifies the character expression or memo field from which the character string is returned.
nStartPosition: Specifies the position in the character expression or memo field cExpression from where the character string is returned. The first character of cExpression is position 1.
                If TALK is set to ON and nStartPosition is greater than the number of characters in cExpression, Visual FoxPro generates an error message. If TALK is set to OFF, the empty string is returned.
nStrReturned:   Specifies the number of characters to return from cExpression. If you omit nStrReturned, characters are returned until the end of the character expression is reached.

*/

func SubStrC(cAText string, nStart, nLenght int) string {

	nEnd := nStart + nLenght - 1
	nLenght = len([]rune(cAText))

	if nEnd > nLenght {
		nEnd = nLenght
	}

	return string([]rune(cAText)[nStart-1 : nEnd])
}

/*
ChrTran Replaces each character in a character expression that matches a character in a second character expression
with the corresponding character in a third character expression.

ChrTran(cSearchedExpression, cSearchExpression, cReplacementExpression)

Parameters
cSearchedExpression:    Specifies the expression in which ChrTran( ) replaces characters.
cSearchExpression:      Specifies the expression containing the characters ChrTran( ) looks for in cSearchedExpression.
cReplacementExpression: Specifies the expression containing the replacement characters.

If a character in cSearchExpression is found in cSearchedExpression, the character in cSearchedExpression is replaced
by a character from cReplacementExpression that's in the same position in cReplacementExpression as the respective character
in cSearchExpression.

If cReplacementExpression has fewer characters than cSearchExpression, the additional characters in cSearchExpression
are deleted from cSearchedExpression. If cReplacementExpression has more characters than cSearchExpression,
the additional characters in cReplacementExpression are ignored.

Remarks: ChrTran( ) translates the character expression cSearchedExpression using the translation expressions cSearchExpression
and cReplacementExpression and returns the resulting character string.

Example: ChrTran("ABCD", "ABC", "YZChrTran)  // ReturnStr: YZD
*/

func ChrTran(sAText, sSearch, sReplace string) string {

	var nRepLen int
	var cSearch, cReplace string

	nRepLen = len(sReplace)

	for ii, rSearch := range sSearch {

		cSearch = string(rSearch)

		if ii >= nRepLen {
			cReplace = ""
		} else {
			cReplace = string(sReplace[ii])
		}

		sAText = strings.Replace(sAText, cSearch, cReplace, -1)
	}
	return sAText
}

func Str(znum float64, zlen int, zdecimal int) string {
	return strconv.FormatFloat(znum, 'f', zdecimal, 64)[0:zlen]
}

func Transform(zExprArgs interface{}) string { return toString(reflect.ValueOf(zExprArgs)) }

func toString(val reflect.Value) string {

	fmt.Println(val)
	fmt.Println(val.Type)
	fmt.Println(val.Kind)

	var str string
	if !val.IsValid() {
		return "<zero Value>"
	}
	typ := val.Type()
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'g', -1, 64)
	case reflect.Complex64, reflect.Complex128:
		c := val.Complex()
		return strconv.FormatFloat(real(c), 'g', -1, 64) + "+" + strconv.FormatFloat(imag(c), 'g', -1, 64) + "i"
	case reflect.String:
		return val.String()
	case reflect.Bool:
		if val.Bool() {
			return "true"
		} else {
			return "false"
		}
	case reflect.Ptr:
		v := val
		str = typ.String() + "("
		if v.IsNil() {
			str += "0"
		} else {
			str += "&" + toString(v.Elem())
		}
		str += ")"
		return str
	case reflect.Array, reflect.Slice:
		v := val
		str += typ.String()
		str += "{"
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				str += ", "
			}
			str += toString(v.Index(i))
		}
		str += "}"
		return str
	case reflect.Map:
		t := typ
		str = t.String()
		str += "{"
		str += "<can't iterate on maps>"
		str += "}"
		return str
	case reflect.Chan:
		str = typ.String()
		return str
	case reflect.Struct:
		m, ok := val.Type().MethodByName("String")
		if ok {
			str = m.Func.Call([]reflect.Value{val})[0].String()
		} else {

			t := typ
			v := val
			str += t.String()
			str += "{"
			for i, n := 0, v.NumField(); i < n; i++ {
				if i > 0 {
					str += ", "
				}
				str += toString(v.Field(i))
			}
			str += "}"
		}
		return str
	case reflect.Interface:
		return typ.String() + "(" + toString(val.Elem()) + ")"
	case reflect.Func:
		v := val
		return typ.String() + "(" + strconv.FormatUint(uint64(v.Pointer()), 10) + ")"
	default:
		panic("Transform: can't print type " + typ.String())
	}
	return "Transform: can't happen"
}

/* Upper() returns the specified character expression in uppercase.
Usage:   Upper(cExpression)
Params:  cExpression: Specifies the character expression Upper() converts to uppercase.
Remarks: Each lowercase letter (a – z) in the character expression is converted to uppercase (A – Z) in the returned string.
         This is true for high-ASCII characters as well, even if the current font does not display them as lowercase accented characters.
		 Türkçe karakter eklenmesiyle alakalı https://github.com/golang/go/commit/4e2b7f8f41a8ab58489354ff0e2c10a867a4a354
Example: Upper("abcşğıiçüöæä本")  -> "ABCŞĞIIÇÜÖÆÄ本"
*/

func Upper(sAtext string) string   { return strings.ToUpper(sAtext) }
func UpperTr(sAtext string) string { return strings.ToUpperSpecial(unicode.TurkishCase, sAtext) }

/* Lower() returns the specified character expression in uppercase.
Usage:   Lower(cExpression)
Params:  cExpression: Specifies the character expression Lower() converts to uppercase.
Remarks: Each lowercase letter (A – Z) in the character expression is converted to uppercase (a – z) in the returned string.
         This is true for high-ASCII characters as well, even if the current font does not display them as lowercase accented characters.
Example: Lower("ABCŞĞİIÇÜÖÆÄ本")  -> "abcşğiiçüöæä本"

*/

func Lower(sAtext string) string   { return strings.ToLower(sAtext) }
func LowerTr(sAtext string) string { return strings.ToLowerSpecial(unicode.TurkishCase, sAtext) }
