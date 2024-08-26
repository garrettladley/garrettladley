package types

import (
	"strings"
	"time"

	"github.com/garrettladley/left-arrow/subs/reports/pkg/dts/constants"
)

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(constants.YYYY_MM_DD, s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) Into() time.Time {
	return time.Time(d)
}
