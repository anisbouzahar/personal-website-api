package http

import (
	"github.com/anisbouzahar/portfolio-api/internal/app/api/v1"
)

type Server struct {
	Api *v1.API
}

func NewServer(api *v1.API) *Server {

	return &Server{
		Api: api,
	}
}
