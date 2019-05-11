package db

import (
	"errors"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db/lock"
)

//go:generate counterfeiter . ResourceCheckFactory

type ResourceCheckFactory interface {
	ResourceChecks() ([]ResourceCheck, error)
	CreateResourceCheck(int, string) (ResourceCheck, error)
	CreateResourceCheckFromVersion(int, string, atc.Version) (ResourceCheck, error)
}

type resourceCheckFactory struct {
	conn        Conn
	lockFactory lock.LockFactory
}

func NewResourceCheckFactory(conn Conn, lockFactory lock.LockFactory) ResourceCheckFactory {
	return &resourceCheckFactory{
		conn:        conn,
		lockFactory: lockFactory,
	}
}

func (r *resourceCheckFactory) ResourceChecks() ([]ResourceCheck, error) {
	rows, err := resourceChecksQuery.
		OrderBy("r.id ASC").
		RunWith(r.conn).
		Query()
	if err != nil {
		return nil, err
	}

	var resourceChecks []ResourceCheck

	for rows.Next() {
		resourceCheck := &resourceCheck{conn: r.conn, lockFactory: r.lockFactory}

		err := scanResourceCheck(resourceCheck, rows)
		if err != nil {
			return nil, err
		}

		resourceChecks = append(resourceChecks, resourceCheck)
	}

	return resourceChecks, nil
}

func (r *resourceCheckFactory) CreateResourceCheck(reasourceID int, checkType string) (ResourceCheck, error) {
	return nil, errors.New("nope")
}

func (r *resourceCheckFactory) CreateResourceCheckFromVersion(reasourceID int, checkType string, fromVersion atc.Version) (ResourceCheck, error) {
	return nil, errors.New("nope")
}
