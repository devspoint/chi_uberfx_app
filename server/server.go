package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
	"net/http"
)

type (
	HandlerIn struct {
		fx.In

		Handler []Handler `group:"handlers"`
	}

	HandlerOutput struct {
		fx.Out

		Handler Handler `group:"handlers"`
	}

	Handler interface {
		Method() string
		Pattern() string
		http.Handler
	}
)

func Serve(h HandlerIn) {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	for _, handler := range h.Handler {
		r.Method(handler.Method(), handler.Pattern(), handler)
	}

	http.ListenAndServe(":8080", r)
}