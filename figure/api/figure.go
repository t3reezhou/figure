package api

import (
	"fmt"

	"github.com/denominatorzero/web"
	"github.com/denominatorzero/web/inject"
	"github.com/t3reezhou/figure/figure/bll"
)

type FigureHandler struct{}

func NewFigureHandler() *FigureHandler {
	return &FigureHandler{}
}

func (h *FigureHandler) CreateFigure(c *web.AppContext) (map[string]interface{}, error) {
	params := struct {
		Id        int64  `web:"id,required"`
		CompanyId int64  `web:"companyid,required"`
		Name      string `web:"name,required"`
	}{}
	fmt.Printf(">>>>>>>>>>>>>%+v", c.Values)
	if err := inject.Inject(c.Values, &params, "web"); err != nil {
		return nil, err
	}

	if err := bll.FigureBllManager.CreateFigure(params.Id, params.CompanyId, params.Name); err != nil {
		return nil, err
	}
	return map[string]interface{}{"result": "ok"}, nil
}

//
// func (h *FigureHandler) CreateFigure(params param.FigureParams) interface{} {
// 	err := bll.FigureBllManager.CreateFigure(params.Creator, params.CompanyID, params.Name)
// 	if err != nil {
// 		return err
// 	}
// 	return status.OK
// }
// func (h *FigureHandler) GetFigures() interface{} {
// 	figures, err := bll.FigureBllManager.GetFigures()
// 	if err != nil {
// 		return err
// 	}
// 	return map[string]interface{}{"figures": figures}
// }
//
// func (h *FigureHandler) GetFigureByID(params martini.Params) interface{} {
// 	if _, ok := params["id"]; !ok {
// 		return errors.ErrInvalidArgument
// 	}
// 	id, err := strconv.ParseInt(params["id"], 10, 64)
// 	if err != nil {
// 		return err
// 	}
// 	figure, err := bll.FigureBllManager.GetFigure(id)
// 	if err != nil {
// 		return err
// 	}
// 	return figure
// }
