package hello

import (
	"devspoint/chi_uberfx_app/server"
	"go.uber.org/fx"
	"net/http"
)

type Handler struct {}

var HelloModule = fx.Provide(NewHelloHandler)

func NewHelloHandler() server.HandlerOutput {
	return server.HandlerOutput{Handler: Handler{}}
}

func (h Handler) Method() string {
	return http.MethodGet
}

func (h Handler) Pattern() string {
	return "/"
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Chi!"))
}
