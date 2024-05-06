package other

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	// Create a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	// Perform some operation within the timeout context
	select {
	case <-time.After(5 * time.Second): // Simulate a long-running process
		fmt.Fprintf(w, "Handler 1: Hello, World!")
	case <-ctx.Done():
		// If the context times out, respond with a timeout message
		http.Error(w, "Handler 1: Request Timeout", http.StatusRequestTimeout)
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {
	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Perform some operation within the timeout context
	select {
	default: // Simulate a long-running process
		fmt.Fprintf(w, timeouts())

	case <-ctx.Done():
		// If the context times out, respond with a timeout message
		http.Error(w, "Handler 2: Request Timeout", http.StatusRequestTimeout)
	}
}

func timeouts() string {
	time.Sleep(3 * time.Second)
	return "Handler 2: Hello, World!"
}
func TimeOutHTTP() {
	// Create a new ServeMux and register the handler functions
	mux := http.NewServeMux()
	mux.HandleFunc("/handler1", handler1)
	mux.HandleFunc("/handler2", handler2)

	// Create a server and start listening
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("Server listening on port 8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
