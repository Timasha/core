package postgresql

import "github.com/Timasha/core/components"

type Config struct {
	components.Config
	
	Login    string
	Password string

	Address string
	Port    string
	DBName  string
	SSLMode bool

	MaxOpenConns int `default:"10"`
	MaxIdleConns int `default:"10"`
}

func (c *Config) sslMode() string {
	if c.SSLMode {
		return "enable"
	} else {
		return "disable"
	}
}
