package expconf

import (
	"time"
)

// Duration is a wrapper around time.Duration for MarshalJSON and schema.Merge.
type Duration time.Duration

// Merge implements schemas.Mergeable, it will always return our duration without change.
func (l Duration) Merge(othr Duration) Duration {
	return l
}

// MarshalJSON marshals makes Duration transparent to marshaling.
func (l Duration) MarshalJSON() ([]byte, error) {
	return []byte(time.Duration(l).Round(time.Second).String()), nil
}

// UnmarshalJSON marshals makes Duration transparent to unmarshaling.
func (l *Duration) UnmarshalJSON(bytes []byte) error {
	d, err := time.ParseDuration(string(bytes))
	if err != nil {
		return err
	}
	*l = Duration(d)
	return nil
}
