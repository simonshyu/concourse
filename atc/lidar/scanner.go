package lidar

import (
	"context"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/db"
)

func NewScanner(
	logger lager.Logger,
	checkFactory db.CheckFactory,
	secrets creds.Secrets,
) *scanner {
	return &scanner{
		logger:       logger,
		checkFactory: checkFactory,
		secrets:      secrets,
	}

}

type scanner struct {
	logger lager.Logger

	checkFactory db.CheckFactory
	secrets      creds.Secrets
}

func (s *scanner) Run(ctx context.Context) error {

	// fetch all resources

	resources, err := s.checkFactory.Resources()
	if err != nil {
		s.logger.Error("failed-to-get-resources", err)
		return err
	}

	for _, resource := range resources {

		variables := creds.NewVariables(s.secrets, resource.PipelineName(), resource.TeamName())

		resourceTypes, err := s.dbPipeline.ResourceTypes()
		if err != nil {
			s.logger.Error("failed-to-get-resource-types", err)
			return err
		}

		source, err := creds.NewSource(variables, resource.Source()).Evaluate()
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
		resourceConfigScope, err := resource.SetResourceConfig(source, versionedResourceTypes)
		if err != nil {
			s.logger.Error("failed-to-update-resource-config", err)
			return err
		}

		var plan atc.Plan

		// TODO construct a check plan

		err = s.checkFactory.CreateCheck(resourceConfigScope.ID(), plan)
		if err != nil {
			s.logger.Error("failed-to-create-check", err)
			return err
		}

	}
	return nil

}
