package model

type Figure struct {
	Pid       int64  `db:"productionid" json:"pid"`
	Name      string `db:"name"         json:"name"`
	Creator   int64  `db:"creator"      json:"creator"`
	Companyid int64  `db:"companyid"    json:"companyid"`
	Otime     int64  `db:"otime"        json:"otime"`
	Ctime     int64  `db:"ctime"        json:"ctime"`
	Mtime     int64  `db:"mtime"        json:"mtime"`
}
