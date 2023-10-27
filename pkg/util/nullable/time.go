package nullable

import (
	"database/sql"
	"encoding/json"
	"time"
)

type CustomNullTime struct {
	sql.NullTime
}

// func (nt *CustomNullTime) Scan(value interface{}) error {
// 	nt.Time, nt.Valid = value.(time.Time), false
// 	return nil
// }

// func (nt CustomNullTime) Value() (interface{}, error) {
// 	if !nt.Valid {
// 		return nil, nil
// 	}
// 	return nt.Time, nil
// }

func (nt CustomNullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.Time)
}

func (nt *CustomNullTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nt.Valid = false
		return nil
	}
	var t time.Time
	err := json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	nt.Time = t
	nt.Valid = true
	return nil
}
