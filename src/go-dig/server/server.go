package server

import "go-practice/src/go-dig/config"

type Server struct {
	config *config.Config
}

func New(conf *config.Config) *Server {
	return &Server{
		config: conf,
	}
}
