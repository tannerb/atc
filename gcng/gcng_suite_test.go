package gcng_test

import (
	"os"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/concourse/atc/dbng"
	"github.com/concourse/atc/postgresrunner"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/ifrit"

	"testing"
)

func TestGcng(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gcng Suite")
}

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
var (
	postgresRunner postgresrunner.Runner
	dbProcess      ifrit.Process

	dbConn                dbng.Conn
	err                   error
	resourceCacheFactory  dbng.ResourceCacheFactory
	resourceConfigFactory dbng.ResourceConfigFactory

	teamFactory     dbng.TeamFactory
	buildFactory    *dbng.BuildFactory
	pipelineFactory *dbng.PipelineFactory
	resourceFactory *dbng.ResourceFactory

	defaultTeam     *dbng.Team
	defaultPipeline *dbng.Pipeline
	defaultBuild    *dbng.Build

	usedResource *dbng.Resource
)

var _ = BeforeSuite(func() {
	postgresRunner = postgresrunner.Runner{
		Port: 5433 + GinkgoParallelNode(),
	}

	dbProcess = ifrit.Invoke(postgresRunner)

	postgresRunner.CreateTestDB()
})

var _ = BeforeEach(func() {
	postgresRunner.Truncate()

	dbConn = dbng.Wrap(postgresRunner.Open())

	teamFactory = dbng.NewTeamFactory(dbConn)
	buildFactory = dbng.NewBuildFactory(dbConn)
	pipelineFactory = dbng.NewPipelineFactory(dbConn)
	resourceFactory = dbng.NewResourceFactory(dbConn)

	defaultTeam, err = teamFactory.CreateTeam("default-team")
	Expect(err).NotTo(HaveOccurred())

	defaultBuild, err = buildFactory.CreateOneOffBuild(defaultTeam)
	Expect(err).NotTo(HaveOccurred())

	defaultPipeline, err = pipelineFactory.CreatePipeline(defaultTeam, "default-pipeline", "some-config")
	Expect(err).NotTo(HaveOccurred())

	usedResource, err = resourceFactory.CreateResource(
		defaultPipeline,
		"some-resource",
		`{"some":"source"}`,
	)
	Expect(err).NotTo(HaveOccurred())

	setupTx, err := dbConn.Begin()
	Expect(err).ToNot(HaveOccurred())

	baseResourceType := dbng.BaseResourceType{
		Name: "some-base-type",
	}
	_, err = baseResourceType.FindOrCreate(setupTx)
	Expect(err).NotTo(HaveOccurred())

	Expect(setupTx.Commit()).To(Succeed())

	resourceCacheFactory = dbng.NewResourceCacheFactory(dbConn)
	resourceConfigFactory = dbng.NewResourceConfigFactory(dbConn)
})

var _ = AfterSuite(func() {
	dbProcess.Signal(os.Interrupt)
	Eventually(dbProcess.Wait(), 10*time.Second).Should(Receive())
})