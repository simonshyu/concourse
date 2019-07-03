package k8s_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = FDescribe("TSA Service Node Port", func() {

	var (
    helmArgs     []string
	)

	JustBeforeEach(func() {
    setReleaseNameAndNamespace("tnp")

		By("creating a web only deployment in one namespace")
		helmArgs = []string{
			"--set=web.service.type=NodePort",
      "--set=web.service.tsaNodePort=32222",
		}
	})

  It("deployment succeeds", func() {
    deployConcourseChart(releaseName+"-web", helmArgs...)
  })

	AfterEach(func() {
		cleanup(releaseName+"-web", namespace+"-web", nil)
	})

})
