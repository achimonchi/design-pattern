package factory

import (
	"log"
	"net/http"
)

type RouterNative struct {
	host string
	port string
}

func NewRouterNative(host, port string) Router {
	return &RouterNative{
		host: host,
		port: port,
	}
}

func (n *RouterNative) Start() {
	log.Println("Server [NATIVE] running at port ", n.port)
	http.ListenAndServe(n.port, nil)
}
