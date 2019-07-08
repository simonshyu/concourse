package lidar

import (
	"context"
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/db"
	"github.com/pkg/errors"
)

//go:generate counterfeiter . Checkable

type Checkable interface {
	Name() string
	Type() string
	PipelineID() int
	Source() atc.Source
	Tags() atc.Tags
	CheckEvery() string
	CheckTimeout() string
	LastCheckEndTime() time.Time

	SetResourceConfig(
		atc.Source,
		atc.VersionedResourceTypes,
	) (db.ResourceConfigScope, error)

	SetCheckSetupError(error) error
}

func NewScanner(
	logger lager.Logger,
	checkFactory db.CheckFactory,
	planFactory atc.PlanFactory,
	secrets creds.Secrets,
	defaultCheckTimeout time.Duration,
	defaultCheckInterval time.Duration,
) *scanner {
	return &scanner{
		logger:               logger,
		checkFactory:         checkFactory,
		planFactory:          planFactory,
		secrets:              secrets,
		defaultCheckTimeout:  defaultCheckTimeout,
		defaultCheckInterval: defaultCheckInterval,
	}
}

type scanner struct {
	logger lager.Logger

	checkFactory         db.CheckFactory
	planFactory          atc.PlanFactory
	secrets              creds.Secrets
	defaultCheckTimeout  time.Duration
	defaultCheckInterval time.Duration
}

func (s *scanner) Run(ctx context.Context) error {

	lock, acquired, err := s.checkFactory.AcquireScanningLock(s.logger)
	if err != nil {
		s.logger.Error("failed-to-get-scanning-lock", err)
		return err
	}

	if !acquired {
		s.logger.Debug("scanning-already-in-progress")
		return nil
	}

	defer lock.Release()

	resources, err := s.checkFactory.Resources()
	if err != nil {
		s.logger.Error("failed-to-get-resources", err)
		return err
	}

	resourceTypes, err := s.checkFactory.ResourceTypes()
	if err != nil {
		s.logger.Error("failed-to-get-resources", err)
		return err
	}

	waitGroup := new(sync.WaitGroup)

	for _, resource := range resources {
		waitGroup.Add(1)
		go s.scan(waitGroup, resource, resourceTypes)
	}

	waitGroup.Wait()

	return nil
}

func (s *scanner) scan(waitGroup *sync.WaitGroup, resource db.Resource, resourceTypes db.ResourceTypes) {
	defer waitGroup.Done()

	variables := creds.NewVariables(
		s.secrets,
		resource.TeamName(),
		resource.PipelineName(),
	)

	filteredTypes := resourceTypes.BuildTree(resource.Type())

	if err := s.tryCreateCheck(resource, variables, filteredTypes); err != nil {
		s.logger.Error("failed-to-create-check", err)
		s.setCheckError(s.logger, resource, err)
	}
}

func (s *scanner) tryCreateCheck(checkable Checkable, variables creds.Variables, resourceTypes db.ResourceTypes) error {

	var err error

	parentType, found := s.parentType(checkable, resourceTypes)
	if found {
		if err := s.tryCreateCheck(parentType, variables, resourceTypes); err != nil {
			s.logger.Error("failed-to-create-type-check", err)
			s.setCheckError(s.logger, parentType, err)
			return errors.Wrapf(err, "parent type '%v' error", parentType.Name())
		}

		if parentType.Version() == nil {
			return errors.New("parent type has no version")
		}
	}

	timeout := s.defaultCheckTimeout
	if to := checkable.CheckTimeout(); to != "" {
		timeout, err = time.ParseDuration(to)
		if err != nil {
			return err
		}
	}

	interval := s.defaultCheckInterval
	if every := checkable.CheckEvery(); every != "" {
		interval, err = time.ParseDuration(every)
		if err != nil {
			return err
		}
	}

	if time.Now().Before(checkable.LastCheckEndTime().Add(interval)) {
		s.logger.Debug("interval-not-reached", lager.Data{
			"interval": interval,
		})
		return nil
	}

	source, err := creds.NewSource(variables, checkable.Source()).Evaluate()
	if err != nil {
		s.logger.Error("failed-to-evaluate-source", err)
		return err
	}

	versionedResourceTypes, err := creds.NewVersionedResourceTypes(variables, resourceTypes.Deserialize()).Evaluate()
	if err != nil {
		s.logger.Error("failed-to-evaluate-resource-types", err)
		return err
	}

	// This could have changed based on new variable interpolation so update it
	resourceConfigScope, err := checkable.SetResourceConfig(source, versionedResourceTypes)
	if err != nil {
		s.logger.Error("failed-to-update-resource-config", err)
		return err
	}

	fromVersion := make(atc.Version)
	rcv, found, err := resourceConfigScope.LatestVersion()
	if err != nil {
		s.logger.Error("failed-to-get-current-version", err)
		return err
	}

	if found {
		fromVersion = atc.Version(rcv.Version())
	}

	plan := s.planFactory.NewPlan(atc.CheckPlan{
		Name:        checkable.Name(),
		Type:        checkable.Type(),
		Source:      source,
		Tags:        checkable.Tags(),
		Timeout:     timeout.String(),
		FromVersion: fromVersion,

		VersionedResourceTypes: versionedResourceTypes,
	})

	_, created, err := s.checkFactory.CreateCheck(
		resourceConfigScope.ID(),
		resourceConfigScope.ResourceConfig().ID(),
		resourceConfigScope.ResourceConfig().OriginBaseResourceType().ID,
		plan,
	)
	if err != nil {
		s.logger.Error("failed-to-create-check", err)
		return err
	}

	if !created {
		s.logger.Info("check-already-exists")
	}

	return nil
}

func (s *scanner) parentType(checkable Checkable, resourceTypes []db.ResourceType) (db.ResourceType, bool) {
	for _, resourceType := range resourceTypes {
		if resourceType.Name() == checkable.Type() && resourceType.PipelineID() == checkable.PipelineID() {
			return resourceType, true
		}
	}
	return nil, false
}

func (s *scanner) setCheckError(logger lager.Logger, checkable Checkable, err error) {
	setErr := checkable.SetCheckSetupError(err)
	if setErr != nil {
		logger.Error("failed-to-set-check-error", setErr)
	}
}
