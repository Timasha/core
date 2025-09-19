package tracer

import (
	"context"
	"runtime"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func NewSpan(ctx context.Context) (resultCtx context.Context, span trace.Span, caller string) {
	pc, _, _, ok := runtime.Caller(1)

	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		detailsItems := strings.Split(details.Name(), "/")

		if len(detailsItems) != 0 {
			caller = detailsItems[len(detailsItems)-1]
		} else {
			caller = details.Name()
		}
	} else {
		caller = "anonymous"
	}
	resultCtx, span = otel.Tracer("").Start(ctx, caller)

	return resultCtx, span, caller
}
