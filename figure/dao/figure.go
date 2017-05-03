package dao

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/t3reezhou/figure/figure/model"
)

type FigureDao struct {
	db *sqlx.DB
}

const (
	DefaultFigureTable = "figure_figure"
)

func NewFigureDaoManager(db *sqlx.DB) *FigureDao {
	return &FigureDao{db}
}

func (d *FigureDao) CreateTable() error {
	if err := d.CreateFigureTable(); err != nil {
		return err
	}
	return nil
}

func (d *FigureDao) CreateFigureTable() error {
	schema := `
	CREATE TABLE IF NOT EXISTS %s (
  productionid BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(30) NOT NULL,
  creator BIGINT(20) UNSIGNED NOT NULL,
  companyid BIGINT(20) UNSIGNED NOT NULL,
  otime INT(11) UNSIGNED DEFAULT '1363033208',
  ctime INT(11) UNSIGNED DEFAULT '1363033208',
  mtime INT(11) UNSIGNED DEFAULT '1363033208',
  PRIMARY KEY(productionid),
  KEY index_creator (creator),
  KEY index_companyid (companyid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci
`
	_, err := d.db.Exec(fmt.Sprintf(schema, DefaultFigureTable))
	return err
}

func (d *FigureDao) CreateFigure(creator, companyid int64, name string) error {
	now := time.Now().Unix()
	sql := fmt.Sprintf("INSERT INTO %s (name,creator,companyid,otime,ctime,mtime) VALUES('%s', %d, %d, %d, %d, %d)",
		DefaultFigureTable, name, creator, companyid, now, now, now)
	_, err := d.db.Exec(sql)
	return err
}

func (d *FigureDao) GetFigure(productionid int64) (*model.Figure, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE productionid=%d", DefaultFigureTable, productionid)
	var figure model.Figure
	err := d.db.Get(&figure, sql)
	if err != nil {
		return nil, err
	}
	return &figure, nil
}

func (d *FigureDao) GetFigures() ([]*model.Figure, error) {
	sql := fmt.Sprintf("SELECT * FROM %s", DefaultFigureTable)
	var figures []*model.Figure
	err := d.db.Select(&figures, sql)
	if err != nil {
		return nil, err
	}
	return figures, nil
}
