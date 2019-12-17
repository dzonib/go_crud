package handlers

import "github.com/go-chi/chi"

// You dont have to import it cus it's from the same package
func (s *Server) setupEndpoints(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {

		})
	})
}
