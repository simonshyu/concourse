package resourceserver

import (
	"errors"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/db"
)

//go:generate counterfeiter . Checker

type Checker interface {
	Check(checkable db.Checkable, resourceTypes db.ResourceTypes, fromVersion atc.Version) (bool, error)
}

func NewChecker(
	logger lager.Logger,
	secrets creds.Secrets,
	checkFactory db.CheckFactory,
	defaultCheckTimeout time.Duration,
) *checker {
	return &checker{
		logger:              logger,
		secrets:             secrets,
		checkFactory:        checkFactory,
		defaultCheckTimeout: defaultCheckTimeout,
	}
}

type checker struct {
	logger              lager.Logger
	secrets             creds.Secrets
	checkFactory        db.CheckFactory
	defaultCheckTimeout time.Duration
}

func (s *checker) Check(checkable db.Checkable, resourceTypes db.ResourceTypes, fromVersion atc.Version) (bool, error) {

	var err error

	filteredTypes := resourceTypes.Filter(checkable.Type())

	parentType, found := s.parentType(checkable, filteredTypes)
	if found {
		if parentType.Version() == nil {
			return false, errors.New("parent type has no version")
		}
	}

	timeout := s.defaultCheckTimeout
	if to := checkable.CheckTimeout(); to != "" {
		timeout, err = time.ParseDuration(to)
		if err != nil {
			s.logger.Error("failed-to-parse-check-timeout", err)
			return false, err
		}
	}

	variables := creds.NewVariables(
		s.secrets,
		checkable.TeamName(),
		checkable.PipelineName(),
	)

	source, err := creds.NewSource(variables, checkable.Source()).Evaluate()
	if err != nil {
		s.logger.Error("failed-to-evaluate-source", err)
		return false, err
	}

	versionedResourceTypes, err := creds.NewVersionedResourceTypes(variables, filteredTypes.Deserialize()).Evaluate()
	if err != nil {
		s.logger.Error("failed-to-evaluate-resource-types", err)
		return false, err
	}

	// This could have changed based on new variable interpolation so update it
	resourceConfigScope, err := checkable.SetResourceConfig(source, versionedResourceTypes)
	if err != nil {
		s.logger.Error("failed-to-update-resource-config", err)
		return false, err
	}

	if fromVersion == nil {
		rcv, found, err := resourceConfigScope.LatestVersion()
		if err != nil {
			s.logger.Error("failed-to-get-current-version", err)
			return false, err
		}

		if found {
			fromVersion = atc.Version(rcv.Version())
		}
	}

	plan := atc.Plan{
		Check: &atc.CheckPlan{
			Name:        checkable.Name(),
			Type:        checkable.Type(),
			Source:      source,
			Tags:        checkable.Tags(),
			Timeout:     timeout.String(),
			FromVersion: fromVersion,

			VersionedResourceTypes: versionedResourceTypes,
		},
	}

	_, created, err := s.checkFactory.CreateCheck(
		resourceConfigScope.ID(),
		resourceConfigScope.ResourceConfig().ID(),
		resourceConfigScope.ResourceConfig().OriginBaseResourceType().ID,
		plan,
	)
	if err != nil {
		s.logger.Error("failed-to-create-check", err)
		return false, err
	}

	return created, nil
}

func (s *checker) parentType(checkable db.Checkable, resourceTypes []db.ResourceType) (db.ResourceType, bool) {
	for _, resourceType := range resourceTypes {
		if resourceType.Name() == checkable.Type() && resourceType.PipelineID() == checkable.PipelineID() {
			return resourceType, true
		}
	}
	return nil, false
}
