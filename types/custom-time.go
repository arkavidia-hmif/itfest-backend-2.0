package types

import (
	"bytes"
	"gopkg.in/guregu/null.v4"
	"strings"
	"time"
)

type BirthDate null.Time

func (t *BirthDate) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		t.Valid = false
		return nil
	}

	parsedTime, err := time.Parse("2006-01-02", strings.Trim(string(data), "\""))

	if err != nil {
		return err
	}

	t.Time = parsedTime
	t.Valid = true
	return nil
}
