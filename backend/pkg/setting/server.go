package setting

import (
	"gochat/env"
	"time"
)

type Server struct {
	RunMode      string
	EndPoint     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func setupServer(s *Server) {
	if env.IsDevEnv() {
		s.RunMode = "debug"
	} else {
		s.RunMode = "release"
	}

	s.EndPoint = env.GetStr(env.SERVER_ENDPOINT)
	s.ReadTimeout = time.Duration(env.GetInt(env.SERVER_READ_TIMEOUT)) * time.Second
	s.WriteTimeout = time.Duration(env.GetInt(env.SERVER_WRITE_TIMEOUT)) * time.Second
}
