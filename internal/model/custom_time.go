package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

const customTimeFormat = "2006-01-02"

// JSON用のUnmarshal
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == "" {
		ct.Time = time.Time{}
		return nil
	}
	parsedTime, err := time.Parse(customTimeFormat, s)
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}
	ct.Time = parsedTime
	return nil
}

// JSON用のMarshal
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.Time.Format(customTimeFormat) + `"`), nil
}

// GORM用のScan (データベース→Go)
func (ct *CustomTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		ct.Time = v
		return nil
	case []byte:
		parsedTime, err := time.Parse(customTimeFormat, string(v))
		if err != nil {
			return fmt.Errorf("failed to parse date: %v", err)
		}
		ct.Time = parsedTime
		return nil
	case string:
		parsedTime, err := time.Parse(customTimeFormat, v)
		if err != nil {
			return fmt.Errorf("failed to parse date: %v", err)
		}
		ct.Time = parsedTime
		return nil
	default:
		return fmt.Errorf("unsupported scan type for CustomTime: %T", value)
	}
}

// GORM用のValue (Go→データベース)
func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time.Format(customTimeFormat), nil
}

