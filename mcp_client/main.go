package main

import (
	"context"
	"log"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/http"
)

func main() {
	// Create an HTTP client transport
	transport := http.NewHTTPClientTransport("/mcp")
	transport.WithBaseURL("http://localhost:8080")

	// Create client with the HTTP transport
	client := mcp.NewClient(transport)

	// Use the client with context
	ctx := context.Background()
	response, err := client.CallTool(ctx, "hello", map[string]interface{}{
		"name": "World",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response: %v", response)
}
