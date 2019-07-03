package db_test

import (
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Check", func() {
	var (
		err                 error
		check               db.Check
		resourceConfigScope db.ResourceConfigScope
		ubrt                *db.UsedBaseResourceType
	)

	BeforeEach(func() {

		setupTx, err := dbConn.Begin()
		Expect(err).ToNot(HaveOccurred())

		brt := db.BaseResourceType{
			Name: "some-base-resource-type",
		}

		ubrt, err = brt.FindOrCreate(setupTx, false)
		Expect(err).NotTo(HaveOccurred())
		Expect(setupTx.Commit()).To(Succeed())

		resourceConfigScope, err = defaultResource.SetResourceConfig(atc.Source{"some": "repository"}, atc.VersionedResourceTypes{})
		Expect(err).NotTo(HaveOccurred())

	})

	JustBeforeEach(func() {
		check, err = checkFactory.CreateCheck(resourceConfigScope.ID(), ubrt.ID, atc.Plan{})
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("ResourceConfigScopeID", func() {
		Context("when looking up resource config scope succeeds", func() {
			It("returns the resource", func() {
				Expect(check).NotTo(BeNil())
			})
		})
	})
})
