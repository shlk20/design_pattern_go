package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	nginxServer := newNginxServer()
	statusUrl := "/app/status"
	createUserUrl := "/create/user"

	code, message := nginxServer.handleRequest(statusUrl, "GET")
	fmt.Printf("\nUrl: %s\nHttp Code: %d\nMessage: %s\n", statusUrl, code, message)

	code, message = nginxServer.handleRequest(createUserUrl, "POST")
	fmt.Printf("\nUrl: %s\nHttp Code: %d\nMessage: %s\n", createUserUrl, code, message)

	code, message = nginxServer.handleRequest(statusUrl, "POST")
	fmt.Printf("\nUrl: %s\nHttp Code: %d\nMessage: %s\n", statusUrl, code, message)

	code, message = nginxServer.handleRequest(createUserUrl, "POST")
	fmt.Printf("\nUrl: %s\nHttp Code: %d\nMessage: %s\n", createUserUrl, code, message)

	code, message = nginxServer.handleRequest(statusUrl, "GET")
	fmt.Printf("\nUrl: %s\nHttp Code: %d\nMessage: %s\n", statusUrl, code, message)
}
