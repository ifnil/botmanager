package common

import "database/sql"

type DatabaseAction int

const (
	Exec DatabaseAction = iota
	Query
)

type DatabaseRequest struct {
	Args   any
	Stmt   *sql.Stmt
	Action DatabaseAction
}
