package db

import (
	"time"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db/lock"
)

//go:generate counterfeiter . ResourceCheck

type ResourceCheck interface {
	ID() int
	Resource() (Resource, error)
	Start() error
	Timeout() time.Duration
	FromVersion() atc.Version
	Finish() error
	FinishWithError(message string) error
}

const (
	CheckTypeResource     = "resource"
	CheckTypeResourceType = "resource_type"
)

var resourceChecksQuery = psql.Select("r.id, r.resource_config_scope_id, r.start_time, r.end_time, r.timeout, r.from_version, r.check_error, r.create_time").
	From("resource_checks r")

type resourceCheck struct {
	conn        Conn
	lockFactory lock.LockFactory
}

func (r *resourceCheck) ID() int                    { return 0 }
func (r *resourceCheck) ResourceConfigScopeID() int { return 0 }
func (r *resourceCheck) StartTime() time.Time       { return time.Now() }
func (r *resourceCheck) EndTime() time.Time         { return time.Now() }
func (r *resourceCheck) CreateTime() time.Time      { return time.Now() }
func (r *resourceCheck) Timeout() time.Duration     { return time.Second }
func (r *resourceCheck) FromVersion() atc.Version   { return nil }

func (r *resourceCheck) Start() error {
	return nil
}
func (r *resourceCheck) Finish() error {
	return nil
}
func (r *resourceCheck) FinishWithError(message string) error {
	return nil
}
func (r *resourceCheck) Resource() (Resource, error) {
	return nil, nil
}

func scanResourceCheck(r *resourceCheck, row scannable) error {
	return nil
}
