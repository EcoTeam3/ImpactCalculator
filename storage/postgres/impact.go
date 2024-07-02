package postgres

import "database/sql"


type NewImpact struct{
	Db *sql.DB
}

func NewImpactRepo(db *sql.DB)*NewImpact{
	return &NewImpact{Db: db}
}

