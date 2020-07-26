package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprint(time.Time(t).Unix())), nil
}

func (t Time) IsZero() bool {
	return time.Time(t).IsZero()
}

func (t Time) Unix() int64 {
	return time.Time(t).Unix()
}

func Now() Time {
	return Time(time.Now())
}

func NowPtr() *Time {
	now := Now()
	return &now
}

type StringList []string

func (s StringList) Value() (driver.Value, error) {
	return strings.Join(s, ";"), nil
}

func (s *StringList) Scan(value interface{}) error {
	raw := sql.NullString{}
	if err := raw.Scan(value); err != nil {
		return err
	}
	*s = strings.Split(raw.String, ";")
	return nil
}

type HashMap map[string]interface{}

func (m *HashMap) Scan(src interface{}) (err error) {
	var data []byte
	switch s := src.(type) {
	case string:
		data = []byte(s)
	case []byte:
		data = s
	}
	if data != nil && len(data) > 1 {
		err = json.Unmarshal(data, m)
	}
	return
}

func (m HashMap) String() string {
	if data, err := json.Marshal(&m); err == nil {
		return string(data)
	}
	return "{}"
}

func (m HashMap) Value() (driver.Value, error) {
	return m.String(), nil
}

type StringMap map[string]string

func (m *StringMap) Scan(src interface{}) (err error) {
	var data []byte
	switch s := src.(type) {
	case string:
		data = []byte(s)
	case []byte:
		data = s
	}
	if data != nil && len(data) > 1 {
		err = json.Unmarshal(data, m)
	}
	return
}

func (m StringMap) String() string {
	if data, err := json.Marshal(&m); err == nil {
		return string(data)
	}
	return "[]"
}

func (m StringMap) Value() (driver.Value, error) {
	return m.String(), nil
}
