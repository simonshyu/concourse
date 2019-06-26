package lidar_test

import (
	"context"

	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/db/dbfakes"
	"github.com/concourse/concourse/atc/engine"
	"github.com/concourse/concourse/atc/engine/enginefakes"
	"github.com/concourse/concourse/atc/lidar"
)

type Checker interface {
	Run(context.Context) error
}

var _ = Describe("Checker", func() {
	var (
		// err error

		fakeCheckFactory *dbfakes.FakeCheckFactory
		fakeEngine       *enginefakes.FakeEngine

		checker *lidar.Checker
		logger  *lagertest.TestLogger
	)

	BeforeEach(func() {
		fakeCheckFactory = new(dbfakes.FakeCheckFactory)
		fakeEngine = new(enginefakes.FakeEngine)

		logger = lagertest.NewTestLogger("test")
		checker = lidar.NewChecker(
			logger,
			fakeCheckFactory,
			fakeEngine,
		)
	})

	JustBeforeEach(func() {
		checker.Check()
	})

	Describe("Check", func() {
		var inQueueChecks []*dbfakes.FakeCheck
		var engineChecks []*enginefakes.FakeRunnable

		BeforeEach(func() {
			inQueueChecks = []*dbfakes.FakeCheck{
				new(dbfakes.FakeCheck),
				new(dbfakes.FakeCheck),
				new(dbfakes.FakeCheck),
			}
			returnedChecks := []db.Check{
				inQueueChecks[0],
				inQueueChecks[1],
				inQueueChecks[2],
			}

			fakeCheckFactory.ChecksReturns(returnedChecks, nil)

			engineChecks = []*enginefakes.FakeRunnable{}
			fakeEngine.NewCheckStub = func(build db.Check) engine.Runnable {
				engineCheck := new(enginefakes.FakeRunnable)
				engineChecks = append(engineChecks, engineCheck)
				return engineCheck
			}
		})

		It("resumes all currently in-flight builds", func() {
			Eventually(engineChecks[0].RunCallCount).Should(Equal(1))
			Eventually(engineChecks[1].RunCallCount).Should(Equal(1))
			Eventually(engineChecks[2].RunCallCount).Should(Equal(1))
		})
	})
})
