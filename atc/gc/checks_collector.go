package gc

import (
	"context"

	"code.cloudfoundry.org/lager/lagerctx"
	"github.com/concourse/concourse/atc/db"
)

type checkCollector struct {
	checkLifecycle db.CheckLifecycle
}

func NewCheckCollector(checkLifecycle db.CheckLifecycle) *checkCollector {
	return &checkCollector{
		checkLifecycle: checkLifecycle,
	}
}

func (c *checkCollector) Run(ctx context.Context) error {
	logger := lagerctx.FromContext(ctx).Session("check-collector")

	logger.Debug("start")
	defer logger.Debug("done")

	return c.checkLifecycle.RemoveExpiredChecks()
}
