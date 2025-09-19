package duration

import (
	"encoding/json"
	"time"
)

type Seconds struct {
	time.Duration
}

func (s *Seconds) UnmarshalJSON(bytes []byte) (err error) {
	var seconds int64

	err = json.Unmarshal(bytes, &seconds)
	if err != nil {
		return err
	}

	s.Duration = time.Duration(seconds) * time.Second

	return nil
}

type Hours struct {
	time.Duration
}

func (s *Hours) UnmarshalJSON(bytes []byte) (err error) {
	var hours int64

	err = json.Unmarshal(bytes, &hours)
	if err != nil {
		return err
	}

	s.Duration = time.Duration(hours) * time.Hour

	return nil
}
