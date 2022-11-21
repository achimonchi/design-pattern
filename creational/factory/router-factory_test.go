package factory

import "testing"

func TestRouterFactory(t *testing.T) {
	router := NewRouterFactory("localhost", ":4343").CreateRouter(Router_HttpRouter)

	router.Start()
}
