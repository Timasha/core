package log

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
)

var (
	globalLogger = zerolog.New(os.Stdout)
)

func GetGlobalLogger() zerolog.Logger {
	return globalLogger
}

type GlobalLoggerInitter struct {
	cfg    Config
	writer io.Writer
}

func NewGlobalLoggerInitter(cfg Config, writer io.Writer) *GlobalLoggerInitter {
	return &GlobalLoggerInitter{
		cfg:    cfg,
		writer: writer,
	}
}

func (g *GlobalLoggerInitter) Start(ctx context.Context) error {
	globalLogger = zerolog.New(zerolog.MultiLevelWriter(os.Stdout, g.writer))

	return nil
}

func (g *GlobalLoggerInitter) Stop(ctx context.Context) error {
	return nil
}

func (g *GlobalLoggerInitter) GetName() string {
	return "GlobalLoggerInitter"
}

func Info(msg string) {
	globalLogger.Info().Msg(fmt.Sprintf("%s", msg))
}

func Infof(msg string, args ...any) {
	globalLogger.Info().Msgf(fmt.Sprintf("%s", msg), args)
}

func Error(msg string) {
	globalLogger.Error().Msg(fmt.Sprintf("%s", msg))
}

func Errorf(msg string, args ...any) {
	globalLogger.Error().Msgf(fmt.Sprintf("%s", msg), args)
}

func Warning(msg string) {
	globalLogger.Warn().Msg(fmt.Sprintf("%s", msg))
}

func Warningf(msg string, args ...any) {
	globalLogger.Warn().Msgf(fmt.Sprintf("%s", msg), args)
}

func Fatal(msg string) {
	globalLogger.Fatal().Msg(fmt.Sprintf("%s", msg))
}

func Fatalf(msg string, args ...any) {
	globalLogger.Fatal().Msgf(fmt.Sprintf("%s", msg), args)
}
