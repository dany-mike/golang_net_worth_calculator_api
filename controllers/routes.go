package controllers

import "golang_net_worth_calculator_api/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/protected", middlewares.SetMiddlewareAuthentication(s.HomeProtected)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/signup", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
}
