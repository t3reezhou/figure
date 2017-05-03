package figure

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/t3reezhou/figure/figure/app"
	"github.com/t3reezhou/figure/figure/cfg"
)

type Server struct {
	cfg *cfg.Config
	app *app.App
	web *mux.Router
}

func NewServer(cfg *cfg.Config) (*Server, error) {
	var err error
	s := new(Server)
	s.cfg = cfg
	s.app, err = app.NewApp(cfg)
	if err != nil {
		return nil, err
	}
	s.web = mux.NewRouter()
	RegistUrl(s.web)

	return s, nil
}

func (s *Server) Run() {
	log.Fatal(http.ListenAndServe(":8000", s.web))
}

func (s *Server) Close() {
}
