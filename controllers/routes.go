package controllers

import "golang_net_worth_calculator_api/middlewares"

func (s *Server) initializeRoutes() {
	// Test Routes
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/protected", middlewares.SetMiddlewareAuthentication(s.HomeProtected)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Signup Route
	s.Router.HandleFunc("/signup", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")

	// Item Routes
	s.Router.HandleFunc("/items", middlewares.SetMiddlewareJSON(s.CreateItem)).Methods("POST")
	s.Router.HandleFunc("/items", middlewares.SetMiddlewareJSON(s.GetItems)).Methods("GET")
	s.Router.HandleFunc("/items/{id}", middlewares.SetMiddlewareJSON(s.GetItem)).Methods("GET")
	s.Router.HandleFunc("/items/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateItem))).Methods("PUT")
	s.Router.HandleFunc("/items/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteItem)).Methods("DELETE")
}
