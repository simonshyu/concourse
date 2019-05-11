package db_test

import (
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResourceCheck", func() {
	var resourceCheck db.ResourceCheck
	var err error

	BeforeEach(func() {
		resourceCheck, err = resourceCheckFactory.CreateResourceCheck(1, db.CheckTypeResource)
		Expect(err).NotTo(HaveOccurred())
	})

	FDescribe("Resource", func() {
		Context("when looking up resource succeeds", func() {
			It("returns the resource", func() {
				resource, err := resourceCheck.Resource()

				Expect(err).NotTo(HaveOccurred())
				Expect(resource.ID()).To(Equal(1))
				Expect(resource.Name()).To(Equal("some-resource"))
				Expect(resource.Type()).To(Equal("some-base-resource-type"))
				Expect(resource.Source()).To(Equal(atc.Source{"some": "source"}))
			})
		})
	})
})
