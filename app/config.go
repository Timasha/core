package app

import "github.com/Timasha/core/duration"

type Config struct {
	StartTimeout duration.Seconds `default:"10"`
	StopTimeout  duration.Seconds `default:"10"`
}
