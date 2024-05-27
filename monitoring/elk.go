package monitoring

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
	"go.opentelemetry.io/otel/trace/baggage"
)

func initTracer() func() {
	// Create an OTLP exporter to send traces to Elastic APM.
	exporter, err := elasticsearch.NewRawExporter(elasticsearch.Config{
		Endpoint: "<ELASTICSEARCH_ENDPOINT>",
	})
	if err != nil {
		log.Fatalf("Failed to create exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(resource.NewWithAttributes(semconv.ServiceNameKey.String("example"))),
	)

	otel.SetTracerProvider(tp)

	return func() {
		// Shutdown the exporter before the program exits.
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatalf("Failed to shutdown provider: %v", err)
		}
	}
}

func StartMonitoring() {
	// Initialize the tracer provider.
	shutdown := initTracer()
	defer shutdown()

	// Create a new span to represent the operation.
	tracer := otel.GetTracerProvider().Tracer("example/tracer")

	// Create a new context with baggage.
	ctx := baggage.ContextWithValues(context.Background(), label.String("user_id", "123"))

	// Start a new span within the context.
	ctx, span := tracer.Start(ctx, "exampleSpan")
	defer span.End()

	// Simulate some work.
	time.Sleep(100 * time.Millisecond)

	// Example HTTP request.
	req, err := http.NewRequestWithContext(ctx, "GET", "https://example.com", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Make the HTTP request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Print the response status code.
	fmt.Printf("Response status code: %d\n", resp.StatusCode)
}
