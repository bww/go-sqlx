package sqlx

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

var errInvalidType = errors.New("Invalid type")

type Bytes []byte

func (d Bytes) Value() (driver.Value, error) {
	if len(d) == 0 {
		return nil, nil
	} else {
		return []byte(d), nil
	}
}

func (d *Bytes) Scan(src interface{}) error {
	switch c := src.(type) {
	case nil:
		*d = nil
	case []byte:
		*d = Bytes(c)
	case string:
		*d = Bytes(c)
	default:
		return fmt.Errorf("%w: %T", errInvalidType, src)
	}
	return nil
}

type JSON json.RawMessage

func (d JSON) Value() (driver.Value, error) {
	if len(d) == 0 {
		return nil, nil
	} else {
		return []byte(d), nil
	}
}

func (d *JSON) Scan(src interface{}) error {
	switch c := src.(type) {
	case nil:
		*d = nil
	case []byte:
		*d = JSON(c)
	case string:
		*d = JSON(c)
	default:
		return fmt.Errorf("%w: %T", errInvalidType, src)
	}
	return nil
}
