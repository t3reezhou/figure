package figure

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/t3reezhou/figure/figure/api"
	"github.com/t3reezhou/figure/figure/middleware"
)

func RegistUrl(r *mux.Router) {
	middles := []http.Handler{&middleware.LogMiddle{}, &middleware.ParseMiddle{}, &middleware.WriteMiddle{}}
	apiRouter := r.PathPrefix("/api").Subrouter()

	figureHandler := api.NewFigureHandler()
	apiRouter.Handle("/figures", middleware.Middleware(figureHandler.CreateFigure, middles)).
		Methods("POST")
	apiRouter.Handle("/figures", middleware.Middleware(figureHandler.GetFigures, middles)).
		Methods("GET")
	apiRouter.Handle("/figures/{id:[0-9]+}", middleware.Middleware(figureHandler.GetFigureByID, middles)).
		Methods("GET")
	// apiRouter.Handle("/figures/{id:[0-9]+}", middleware.Middleware(figureHandler.UpdateFigure,middles)).
	// Methods("PUT")
	// apiRouter.Handle("/figures/{id:[0-9]+}", middleware.Middleware(figureHandler.DeleteFigure,middles)).
	// Methods("DELETE")
}
