package db

import (
	"encoding/json"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db/lock"
)

//go:generate counterfeiter . CheckFactory

type CheckFactory interface {
	Checks() ([]Check, error)
	CreateCheck(int, atc.Plan) error
	Resources() ([]Resource, error)
}

type checkFactory struct {
	conn        Conn
	lockFactory lock.LockFactory
}

func NewCheckFactory(conn Conn, lockFactory lock.LockFactory) CheckFactory {
	return &checkFactory{
		conn:        conn,
		lockFactory: lockFactory,
	}
}

func (c *checkFactory) Checks() ([]Check, error) {
	rows, err := checksQuery.
		OrderBy("r.id ASC").
		RunWith(c.conn).
		Query()
	if err != nil {
		return nil, err
	}

	var checks []Check

	for rows.Next() {
		check := &check{conn: c.conn, lockFactory: c.lockFactory}

		err := scanCheck(check, rows)
		if err != nil {
			return nil, err
		}

		checks = append(checks, check)
	}

	return checks, nil
}

func (c *checkFactory) CreateCheck(resourceConfigScopeID int, plan atc.Plan) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return err
	}

	defer Rollback(tx)

	planPayload, err := json.Marshal(plan)
	if err != nil {
		return err
	}

	es := c.conn.EncryptionStrategy()
	encryptedPayload, nonce, err := es.Encrypt(planPayload)
	if err != nil {
		return err
	}

	_, err = psql.Insert("checks").
		Columns(
			"resource_config_scope_id",
			"schema",
			"status",
			"plan",
			"nonce",
		).
		Values(
			resourceConfigScopeID,
			schema,
			CheckStatusPending,
			encryptedPayload,
			nonce,
		).
		RunWith(tx).
		Exec()

	return err
}

func (c *checkFactory) Resources() ([]Resource, error) {
	var resources []Resource

	// condition: pipeline is not paused
	rows, err := resourcesQuery.
		LeftJoin("resource_types rt on rt.name = r.type").
		Where(
			sq.Eq{"p.paused": false},
		).
		RunWith(c.conn).
		Query()

	for rows.Next() {
		r := &resource{
			conn:        c.conn,
			lockFactory: c.lockFactory,
		}

		err = scanResource(r, rows)
		if err != nil {
			return nil, err
		}

		// filter out resources by check interval
		// TODO get this default check interval from somewhere
		interval := 10 * time.Second
		if r.CheckEvery() != "" {
			configuredInterval, err := time.ParseDuration(r.CheckEvery())
			if err != nil {
				return nil, err
			}

			interval = configuredInterval
		}

		if time.Now().Before(r.lastCheckEndTime.Add(interval)) {
			continue
		}

		// TODO: filter out resources if parent doesn't have a version

		resources = append(resources, r)
	}

	return resources, nil
}
