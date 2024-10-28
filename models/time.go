package models

import (
    "database/sql/driver"
    "encoding/json"
    "fmt"
    "time"
)

type CustomTime time.Time

const timeFormat = "2006-01-02T15:04"


// NewCustomTime creates a new CustomTime from a time.Time
func NewCustomTime(t time.Time) *CustomTime {
    ct := CustomTime(t)
    return &ct
}

// Now returns a new CustomTime with the current UTC time
func NewCustomTimeNow() *CustomTime {
    return NewCustomTime(time.Now().UTC())
}

// UnmarshalJSON handles the custom date format from frontend
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }

    if s == "" {
        return nil
    }

    // Parse the time in BRT (UTC-3)
    loc, err := time.LoadLocation("America/Sao_Paulo")
    if err != nil {
        return err
    }

    // Parse time assuming it's in BRT
    t, err := time.ParseInLocation(timeFormat, s, loc)
    if err != nil {
        return err
    }

    // Convert to UTC
    *ct = CustomTime(t.UTC())
    return nil
}

// MarshalJSON converts the time to ISO format
func (ct *CustomTime) MarshalJSON() ([]byte, error) {
    t := time.Time(*ct)
    if t.IsZero() {
        return []byte("null"), nil
    }
    // Return in UTC format
    return json.Marshal(t.UTC().Format(time.RFC3339))
}

// Value implements driver.Valuer
func (ct CustomTime) Value() (driver.Value, error) {
    t := time.Time(ct)
    if t.IsZero() {
        return nil, nil
    }
    return t, nil
}

// Scan implements sql.Scanner
func (ct *CustomTime) Scan(value interface{}) error {
    if value == nil {
        *ct = CustomTime(time.Time{})
        return nil
    }

    switch v := value.(type) {
    case time.Time:
        *ct = CustomTime(v)
        return nil
    }
    return fmt.Errorf("cannot scan %T into CustomTime", value)
}
