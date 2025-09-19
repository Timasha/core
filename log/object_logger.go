package log

import "fmt"

type ObjectLogger struct{}

func (o *ObjectLogger) Info(msg string) {
	globalLogger.Info().Msg(fmt.Sprintf("%s\n", msg))
}

func (o *ObjectLogger) Infof(msg string, args ...any) {
	globalLogger.Info().Msgf(fmt.Sprintf("%s\n", msg), args)
}

func (o *ObjectLogger) Error(msg string) {
	globalLogger.Error().Msg(fmt.Sprintf("%s\n", msg))
}

func (o *ObjectLogger) Errorf(msg string, args ...any) {
	globalLogger.Error().Msgf(fmt.Sprintf("%s\n", msg), args)
}

func (o *ObjectLogger) Warning(msg string) {
	globalLogger.Warn().Msg(fmt.Sprintf("%s\n", msg))
}

func (o *ObjectLogger) Warningf(msg string, args ...any) {
	globalLogger.Warn().Msgf(fmt.Sprintf("%s\n", msg), args)
}

func (o *ObjectLogger) Fatal(msg string) {
	globalLogger.Fatal().Msg(fmt.Sprintf("%s\n", msg))
}

func (o *ObjectLogger) Fatalf(msg string, args ...any) {
	globalLogger.Fatal().Msgf(fmt.Sprintf("%s\n", msg), args)
}
