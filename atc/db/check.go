package db

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db/lock"
	"github.com/lib/pq"
)

type CheckStatus string

const (
	CheckStatusPending   CheckStatus = "pending"
	CheckStatusStarted   CheckStatus = "started"
	CheckStatusSucceeded CheckStatus = "succeeded"
	CheckStatusErrored   CheckStatus = "errored"
)

//go:generate counterfeiter . Check

type Check interface {
	ID() int
	ResourceConfigScope() (ResourceConfigScope, error)
	Start() error
	Finish() error
	FinishWithError(err error) error

	Status() CheckStatus
	IsRunning() bool
	AcquireTrackingLock() (lock.Lock, bool, error)
}

const (
	CheckTypeResource     = "resource"
	CheckTypeResourceType = "resource_type"
)

var checksQuery = psql.Select("c.id, c.resource_config_scope_id, c.status, c.schema, c.create_time, c.start_time, c.end_time, c.plan, c.nonce").
	From("checks c")

type check struct {
	id                    int
	resourceConfigScopeID int
	status                CheckStatus

	schema string
	plan   atc.Plan

	createTime time.Time
	startTime  time.Time
	endTime    time.Time

	conn        Conn
	lockFactory lock.LockFactory
}

func (c *check) ID() int                    { return c.id }
func (c *check) ResourceConfigScopeID() int { return c.resourceConfigScopeID }
func (c *check) Status() CheckStatus        { return c.status }
func (c *check) Schema() time.Time          { return c.endTime }
func (c *check) Plan() atc.Plan             { return c.plan }
func (c *check) CreateTime() time.Time      { return c.createTime }
func (c *check) StartTime() time.Time       { return c.startTime }
func (c *check) EndTime() time.Time         { return c.endTime }

func (c *check) Start() error {
	return nil
}

func (c *check) Finish() error {
	return nil
}

func (c *check) FinishWithError(err error) error {
	return nil
}

func (c *check) ResourceConfigScope() (ResourceConfigScope, error) {
	return nil, nil
}

func (c *check) IsRunning() bool {
	return false
}

func (c *check) AcquireTrackingLock() (lock.Lock, bool, error) {
	return nil, false, nil
}

func scanCheck(c *check, row scannable) error {
	var (
		resourceConfigScopeID          sql.NullInt64
		createTime, startTime, endTime pq.NullTime
		schema, plan, nonce            sql.NullString
		status                         string
	)

	err := row.Scan(&c.id, &resourceConfigScopeID, &status, &schema, &createTime, &startTime, &endTime, &plan, &nonce)
	if err != nil {
		return err
	}

	var noncense *string
	if nonce.Valid {
		noncense = &nonce.String
	}

	es := c.conn.EncryptionStrategy()
	decryptedConfig, err := es.Decrypt(string(plan.String), noncense)
	if err != nil {
		return err
	}

	err = json.Unmarshal(decryptedConfig, &c.plan)
	if err != nil {
		return err
	}

	c.status = CheckStatus(status)
	c.schema = schema.String
	c.resourceConfigScopeID = int(resourceConfigScopeID.Int64)
	c.createTime = createTime.Time
	c.startTime = startTime.Time
	c.endTime = endTime.Time

	return nil
}
