package str

import (
	"github.com/elgs/gostrgen"
)

func RanStr(length int) string {
	if length <= 0 {
		length = 6
	}
	charSet := gostrgen.LowerUpper | gostrgen.Digit
	includes := ""   // optionally include some additional letters
	excludes := "Ol" //exclude big 'O' and small 'l' to avoid confusion with zero and one.

	if str, err := gostrgen.RandGen(length, charSet, includes, excludes); err == nil {
		return str
	} else {
		return "错误"
	}

}
