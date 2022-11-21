package factory

type RouterType string

const (
	Router_Native     RouterType = "native"
	Router_HttpRouter RouterType = "httprouter"
	Router_Gin        RouterType = "gin"
)

type RouterFactory struct {
	host string
	port string
}

func NewRouterFactory(host, port string) *RouterFactory {
	return &RouterFactory{
		host: host,
		port: port,
	}
}

func (r *RouterFactory) CreateRouter(routerType RouterType) Router {
	switch routerType {
	case Router_Native:
		return NewRouterNative(r.host, r.port)
	case Router_HttpRouter:
		return NewHttpRouter(r.host, r.port)
	default:
		return nil
	}
}
