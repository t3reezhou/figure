package api

import (
	"net/http"

	"github.com/t3reezhou/figure/figure/bll"
	"github.com/t3reezhou/figure/figure/status"
	"github.com/t3reezhou/figure/figure/status/errors"
	"github.com/t3reezhou/figure/figure/util"
)

type FigureHandler struct{}

func NewFigureHandler() *FigureHandler {
	return &FigureHandler{}
}

func (h *FigureHandler) CreateFigure(rw http.ResponseWriter, r *http.Request) {
	params := struct {
		CompanyId int64  `web:"companyid,required"`
		Name      string `web:"name,required"`
	}{}
	if err := util.Inject(r.Context(), &params); err != nil {
		util.Write(r, err)
		return
	}
	if err := bll.FigureBllManager.CreateFigure(2, params.CompanyId, params.Name); err != nil {
		util.Write(r, err)
		return
	}
	util.Write(r, status.OK)
}

func (h *FigureHandler) GetFigures(rw http.ResponseWriter, r *http.Request) {
	figures, err := bll.FigureBllManager.GetFigures()
	if err != nil {
		util.Write(r, err)
		return
	}
	util.Write(r, figures)
}

func (h *FigureHandler) GetFigureByID(rw http.ResponseWriter, r *http.Request) {
	params := struct {
		Id int64 `web:"id"`
	}{}
	if err := util.Inject(r.Context(), &params); err != nil {
		util.Write(r, err)
		return
	}
	figure, err := bll.FigureBllManager.GetFigure(params.Id)
	if err != nil {
		err = errors.Annotater(err, "context")
		util.Write(r, err)
		return
	}
	util.Write(r, figure)
}
