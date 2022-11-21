package factory

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HttpRouter struct {
	host string
	port string
}

func NewHttpRouter(host, port string) Router {
	return &HttpRouter{
		host: host,
		port: port,
	}
}

func (n *HttpRouter) Start() {
	log.Println("Server [HTTP ROUTER] running at port ", n.port)
	router := httprouter.New()

	http.ListenAndServe(n.port, router)
}
