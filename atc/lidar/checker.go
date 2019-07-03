package lidar

import (
	"context"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/engine"
)

func NewChecker(
	logger lager.Logger,
	checkFactory db.CheckFactory,
	engine engine.Engine,
) *Checker {
	return &Checker{
		logger:       logger,
		checkFactory: checkFactory,
		engine:       engine,
	}
}

type Checker struct {
	logger lager.Logger

	checkFactory db.CheckFactory
	engine       engine.Engine
}

func (c *Checker) Run(ctx context.Context) error {
	cLog := c.logger.Session("check")

	cLog.Debug("start")
	defer cLog.Debug("done")

	checks, err := c.checkFactory.PendingChecks()
	if err != nil {
		c.logger.Error("failed-to-fetch-resource-checks", err)
		return err
	}

	for _, check := range checks {
		btLog := cLog.WithData(lager.Data{
			"check": check.ID(),
		})

		engineCheck := c.engine.NewCheck(check)
		go engineCheck.Run(btLog)
	}

	return nil
}
