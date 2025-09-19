package tracer

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type Tracer struct {
	cfg Config

	cli      otlptrace.Client
	exporter *otlptrace.Exporter
}

func New(cfg Config) *Tracer {
	return &Tracer{cfg: cfg}
}

func (t *Tracer) Start(ctx context.Context) (err error) {
	opts := []otlptracegrpc.Option{otlptracegrpc.WithEndpoint(t.cfg.TracingAddress)}
	if t.cfg.Insecure {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}

	t.cli = otlptracegrpc.NewClient(opts...)

	t.exporter, err = otlptrace.New(context.WithoutCancel(ctx), t.cli)
	if err != nil {
		return err
	}

	rsc, err :=
		resource.Merge(
			resource.Default(),
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(t.cfg.AppName),
			),
		)
	if err != nil {
		return err
	}

	otel.SetTracerProvider(sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(t.exporter),
		sdktrace.WithResource(rsc),
	))

	return nil
}

func (t *Tracer) Stop(ctx context.Context) (err error) {
	err = t.exporter.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tracer) GetName() string {
	return "tracer"
}

func (t *Tracer) IsEnabled() bool {
	return t.cfg.IsEnabled
}
