package tracer

import "github.com/Timasha/core/components"

type Config struct {
	components.Config
	AppName        string
	TracingAddress string
	Insecure       bool `default:"true"`
}
