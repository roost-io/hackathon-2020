package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/label"

	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type stringKey string

const traceBallotIDKey stringKey = "ballot_id"
const traceVoterBoothIDKey stringKey = "booth_id"

// initTracer creates a new trace provider instance and registers it as global trace provider.
func initTracer() func() {
	// Create and install Jaeger export pipeline
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint("http://jaeger-collector:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: "ballot",
			Tags: []label.KeyValue{
				label.String("exporter", "jaeger"),
				label.String("app", "voter"),
				label.String("vendor", "zbio"),
			},
		}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	if err != nil {
		log.Fatal(err)
	}

	return func() {
		flush()
	}
}

func getBallotID() string {
	hostname, _ := os.Hostname()
	return hostname
}
func generateRandomID() uuid.UUID {
	return uuid.New()
}

func newRequestContext(ctx context.Context, r *http.Request) context.Context {
	reqID := r.Header.Get("X-Request-ID")
	if reqID == "" {
		uid := generateRandomID()
		reqID = "" + uid.String()
	}
	ctx = context.WithValue(ctx, traceVoterBoothIDKey, reqID)
	return context.WithValue(ctx, traceBallotIDKey, getBallotID)
}

func middlewareTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := newRequestContext(r.Context(), r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// fn := initTracer()
// defer fn()

// ctx := context.Background()
// ctx.Value("")

// func InitTrace() {
// 	jaegerExporter, err := jaeger.NewRawExporter(
// 		jaeger.WithCollectorEndpoint("http://jaeger-collector:14268/api/traces"),
// 		jaeger.WithProcess(jaeger.Process{
// 			ServiceName: "ballot",
// 			Tags: []core.KeyValue{
// 				key.String("exporter", "jaeger"),
// 				key.String("version", "ballot:v1"),
// 			},
// 		}),
// 	)

// 	if err != nil {
// 		log.Fatalf("Error encountered acquiring Jaeger Exporter: %v", err)
// 	}
// 	tp, err := sdktrace.NewProvider(
// 		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
// 		sdktrace.WithSyncer(jaegerExporter),
// 	)
// 	if err != nil {
// 		log.Fatalf("Error encountered setting sdktrace provider: %v", err)
// 	}
// 	global.SetTraceProvider(tp)
// }
