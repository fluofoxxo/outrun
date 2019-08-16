package responseobjs

import (
	"strconv"
)

type ErrorMessage string

func (em ErrorMessage) MarshalJSON() ([]byte, error) {
	return []byte(strconv.QuoteToASCII(string(em))), nil
}
