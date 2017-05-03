package figure

import (
	"github.com/denominatorzero/web"
	"github.com/denominatorzero/web/middleware"
	"github.com/gorilla/mux"
	"github.com/t3reezhou/figure/figure/api"
)

func RegistUrl(r *mux.Router) {
	figureHandler := api.NewFigureHandler()
	r.Handle("/figure/{id:[0-9]+}", web.NewAppHandler(figureHandler.CreateFigure, middleware.ParseMiddle{})).
		Methods("POST")
}
