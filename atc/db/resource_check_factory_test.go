package db_test

import (
	"github.com/concourse/concourse/atc/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResourceCheckFactory", func() {
	Describe("ResourceChecks", func() {
		Context("when looking up the resource check succeeds", func() {
			var resourceCheck db.ResourceCheck
			var err error

			BeforeEach(func() {
				resourceCheck, err = resourceCheckFactory.CreateResourceCheck(1, db.CheckTypeResource)
				Expect(err).NotTo(HaveOccurred())
			})

			FIt("returns the resource check", func() {
				checks, err := resourceCheckFactory.ResourceChecks()

				Expect(err).NotTo(HaveOccurred())
				Expect(checks).To(HaveLen(1))
				Expect(checks[0]).To(Equal(resourceCheck))
			})
		})
	})
})
