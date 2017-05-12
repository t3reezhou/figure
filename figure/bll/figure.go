package bll

import (
	"github.com/t3reezhou/figure/figure/dao"
	"github.com/t3reezhou/figure/figure/model"
	"github.com/t3reezhou/figure/figure/status/errors"
)

type FigureBll struct{}

func NewFigureBll() *FigureBll {
	return &FigureBll{}
}

func (b *FigureBll) GetFigure(id int64) (*model.Figure, error) {
	figure, err := dao.FigureDaoManager.GetFigure(id)
	if err != nil {
		return nil, errors.ErrFigureNotExist
	}

	return figure, nil
}
func (b *FigureBll) GetFigures() ([]*model.Figure, error) {
	figures, err := dao.FigureDaoManager.GetFigures()
	if err != nil {
		return nil, err
	}

	return figures, nil
}

func (b *FigureBll) CreateFigure(creator, companyid int64, name string) error {
	err := dao.FigureDaoManager.CreateFigure(creator, companyid, name)
	if err != nil {
		return err
	}
	return nil
}
