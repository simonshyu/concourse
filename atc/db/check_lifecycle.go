package db

import (
	sq "github.com/Masterminds/squirrel"
)

//go:generate counterfeiter . CheckLifecycle

type CheckLifecycle interface {
	RemoveExpiredChecks() error
}

type checkLifecycle struct {
	conn Conn
}

func NewCheckLifecycle(conn Conn) *checkLifecycle {
	return &checkLifecycle{
		conn: conn,
	}
}

func (lifecycle *checkLifecycle) RemoveExpiredChecks() error {

	_, err := psql.Delete("checks").
		Where(
			sq.And{
				sq.Expr("create_time < NOW() - interval '24 hours'"),
				sq.NotEq{"status": CheckStatusStarted},
			},
		).
		RunWith(lifecycle.conn).
		Exec()

	return err
}
