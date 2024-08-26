package types

import (
	"strconv"
	"strings"
)

type StringInt int

func (si *StringInt) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*si = StringInt(i)
	return nil
}

func (si StringInt) Into() int {
	return int(si)
}
