package dss

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func NewServer(h *Handler) *Server {
	s := &Server{
		Router: chi.NewRouter(),
	}

	s.SetupRouter(h)
	return s
}

type Server struct {
	Router *chi.Mux
}

func (s *Server) SetupRouter(h *Handler) {
	s.Router.Use(middleware.Logger)
	s.Router.Route("/api/moora", func(r chi.Router) {
		r.Post("/", h.Moora)
	})
}

func (s Server) Serve(addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: s.Router,
	}
	log.Info().Msgf("Starting up server at %+v", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")

	}
}
