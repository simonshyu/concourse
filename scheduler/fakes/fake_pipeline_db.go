// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/config"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/db/algorithm"
	"github.com/concourse/atc/scheduler"
)

type FakePipelineDB struct {
	GetJobStub        func(job string) (db.SavedJob, error)
	getJobMutex       sync.RWMutex
	getJobArgsForCall []struct {
		job string
	}
	getJobReturns struct {
		result1 db.SavedJob
		result2 error
	}
	GetRunningBuildsBySerialGroupStub        func(jobName string, serialGroups []string) ([]db.Build, error)
	getRunningBuildsBySerialGroupMutex       sync.RWMutex
	getRunningBuildsBySerialGroupArgsForCall []struct {
		jobName      string
		serialGroups []string
	}
	getRunningBuildsBySerialGroupReturns struct {
		result1 []db.Build
		result2 error
	}
	GetNextPendingBuildBySerialGroupStub        func(jobName string, serialGroups []string) (db.Build, bool, error)
	getNextPendingBuildBySerialGroupMutex       sync.RWMutex
	getNextPendingBuildBySerialGroupArgsForCall []struct {
		jobName      string
		serialGroups []string
	}
	getNextPendingBuildBySerialGroupReturns struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	UpdateBuildPreparationStub        func(prep db.BuildPreparation) error
	updateBuildPreparationMutex       sync.RWMutex
	updateBuildPreparationArgsForCall []struct {
		prep db.BuildPreparation
	}
	updateBuildPreparationReturns struct {
		result1 error
	}
	IsPausedStub        func() (bool, error)
	isPausedMutex       sync.RWMutex
	isPausedArgsForCall []struct{}
	isPausedReturns     struct {
		result1 bool
		result2 error
	}
	LoadVersionsDBStub        func() (*algorithm.VersionsDB, error)
	loadVersionsDBMutex       sync.RWMutex
	loadVersionsDBArgsForCall []struct{}
	loadVersionsDBReturns     struct {
		result1 *algorithm.VersionsDB
		result2 error
	}
	GetNextInputVersionsStub        func(versions *algorithm.VersionsDB, job string, inputs []config.JobInput) ([]db.BuildInput, bool, error)
	getNextInputVersionsMutex       sync.RWMutex
	getNextInputVersionsArgsForCall []struct {
		versions *algorithm.VersionsDB
		job      string
		inputs   []config.JobInput
	}
	getNextInputVersionsReturns struct {
		result1 []db.BuildInput
		result2 bool
		result3 error
	}
	UseInputsForBuildStub        func(buildID int, inputs []db.BuildInput) error
	useInputsForBuildMutex       sync.RWMutex
	useInputsForBuildArgsForCall []struct {
		buildID int
		inputs  []db.BuildInput
	}
	useInputsForBuildReturns struct {
		result1 error
	}
	CreateJobBuildStub        func(job string) (db.Build, error)
	createJobBuildMutex       sync.RWMutex
	createJobBuildArgsForCall []struct {
		job string
	}
	createJobBuildReturns struct {
		result1 db.Build
		result2 error
	}
	CreateJobBuildForCandidateInputsStub        func(job string) (db.Build, bool, error)
	createJobBuildForCandidateInputsMutex       sync.RWMutex
	createJobBuildForCandidateInputsArgsForCall []struct {
		job string
	}
	createJobBuildForCandidateInputsReturns struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	UpdateBuildToScheduledStub        func(buildID int) (bool, error)
	updateBuildToScheduledMutex       sync.RWMutex
	updateBuildToScheduledArgsForCall []struct {
		buildID int
	}
	updateBuildToScheduledReturns struct {
		result1 bool
		result2 error
	}
	GetJobBuildForInputsStub        func(job string, inputs []db.BuildInput) (db.Build, bool, error)
	getJobBuildForInputsMutex       sync.RWMutex
	getJobBuildForInputsArgsForCall []struct {
		job    string
		inputs []db.BuildInput
	}
	getJobBuildForInputsReturns struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	GetNextPendingBuildStub        func(job string) (db.Build, bool, error)
	getNextPendingBuildMutex       sync.RWMutex
	getNextPendingBuildArgsForCall []struct {
		job string
	}
	getNextPendingBuildReturns struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	SaveResourceVersionsStub        func(atc.ResourceConfig, []atc.Version) error
	saveResourceVersionsMutex       sync.RWMutex
	saveResourceVersionsArgsForCall []struct {
		arg1 atc.ResourceConfig
		arg2 []atc.Version
	}
	saveResourceVersionsReturns struct {
		result1 error
	}
}

func (fake *FakePipelineDB) GetJob(job string) (db.SavedJob, error) {
	fake.getJobMutex.Lock()
	fake.getJobArgsForCall = append(fake.getJobArgsForCall, struct {
		job string
	}{job})
	fake.getJobMutex.Unlock()
	if fake.GetJobStub != nil {
		return fake.GetJobStub(job)
	} else {
		return fake.getJobReturns.result1, fake.getJobReturns.result2
	}
}

func (fake *FakePipelineDB) GetJobCallCount() int {
	fake.getJobMutex.RLock()
	defer fake.getJobMutex.RUnlock()
	return len(fake.getJobArgsForCall)
}

func (fake *FakePipelineDB) GetJobArgsForCall(i int) string {
	fake.getJobMutex.RLock()
	defer fake.getJobMutex.RUnlock()
	return fake.getJobArgsForCall[i].job
}

func (fake *FakePipelineDB) GetJobReturns(result1 db.SavedJob, result2 error) {
	fake.GetJobStub = nil
	fake.getJobReturns = struct {
		result1 db.SavedJob
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineDB) GetRunningBuildsBySerialGroup(jobName string, serialGroups []string) ([]db.Build, error) {
	fake.getRunningBuildsBySerialGroupMutex.Lock()
	fake.getRunningBuildsBySerialGroupArgsForCall = append(fake.getRunningBuildsBySerialGroupArgsForCall, struct {
		jobName      string
		serialGroups []string
	}{jobName, serialGroups})
	fake.getRunningBuildsBySerialGroupMutex.Unlock()
	if fake.GetRunningBuildsBySerialGroupStub != nil {
		return fake.GetRunningBuildsBySerialGroupStub(jobName, serialGroups)
	} else {
		return fake.getRunningBuildsBySerialGroupReturns.result1, fake.getRunningBuildsBySerialGroupReturns.result2
	}
}

func (fake *FakePipelineDB) GetRunningBuildsBySerialGroupCallCount() int {
	fake.getRunningBuildsBySerialGroupMutex.RLock()
	defer fake.getRunningBuildsBySerialGroupMutex.RUnlock()
	return len(fake.getRunningBuildsBySerialGroupArgsForCall)
}

func (fake *FakePipelineDB) GetRunningBuildsBySerialGroupArgsForCall(i int) (string, []string) {
	fake.getRunningBuildsBySerialGroupMutex.RLock()
	defer fake.getRunningBuildsBySerialGroupMutex.RUnlock()
	return fake.getRunningBuildsBySerialGroupArgsForCall[i].jobName, fake.getRunningBuildsBySerialGroupArgsForCall[i].serialGroups
}

func (fake *FakePipelineDB) GetRunningBuildsBySerialGroupReturns(result1 []db.Build, result2 error) {
	fake.GetRunningBuildsBySerialGroupStub = nil
	fake.getRunningBuildsBySerialGroupReturns = struct {
		result1 []db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineDB) GetNextPendingBuildBySerialGroup(jobName string, serialGroups []string) (db.Build, bool, error) {
	fake.getNextPendingBuildBySerialGroupMutex.Lock()
	fake.getNextPendingBuildBySerialGroupArgsForCall = append(fake.getNextPendingBuildBySerialGroupArgsForCall, struct {
		jobName      string
		serialGroups []string
	}{jobName, serialGroups})
	fake.getNextPendingBuildBySerialGroupMutex.Unlock()
	if fake.GetNextPendingBuildBySerialGroupStub != nil {
		return fake.GetNextPendingBuildBySerialGroupStub(jobName, serialGroups)
	} else {
		return fake.getNextPendingBuildBySerialGroupReturns.result1, fake.getNextPendingBuildBySerialGroupReturns.result2, fake.getNextPendingBuildBySerialGroupReturns.result3
	}
}

func (fake *FakePipelineDB) GetNextPendingBuildBySerialGroupCallCount() int {
	fake.getNextPendingBuildBySerialGroupMutex.RLock()
	defer fake.getNextPendingBuildBySerialGroupMutex.RUnlock()
	return len(fake.getNextPendingBuildBySerialGroupArgsForCall)
}

func (fake *FakePipelineDB) GetNextPendingBuildBySerialGroupArgsForCall(i int) (string, []string) {
	fake.getNextPendingBuildBySerialGroupMutex.RLock()
	defer fake.getNextPendingBuildBySerialGroupMutex.RUnlock()
	return fake.getNextPendingBuildBySerialGroupArgsForCall[i].jobName, fake.getNextPendingBuildBySerialGroupArgsForCall[i].serialGroups
}

func (fake *FakePipelineDB) GetNextPendingBuildBySerialGroupReturns(result1 db.Build, result2 bool, result3 error) {
	fake.GetNextPendingBuildBySerialGroupStub = nil
	fake.getNextPendingBuildBySerialGroupReturns = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePipelineDB) UpdateBuildPreparation(prep db.BuildPreparation) error {
	fake.updateBuildPreparationMutex.Lock()
	fake.updateBuildPreparationArgsForCall = append(fake.updateBuildPreparationArgsForCall, struct {
		prep db.BuildPreparation
	}{prep})
	fake.updateBuildPreparationMutex.Unlock()
	if fake.UpdateBuildPreparationStub != nil {
		return fake.UpdateBuildPreparationStub(prep)
	} else {
		return fake.updateBuildPreparationReturns.result1
	}
}

func (fake *FakePipelineDB) UpdateBuildPreparationCallCount() int {
	fake.updateBuildPreparationMutex.RLock()
	defer fake.updateBuildPreparationMutex.RUnlock()
	return len(fake.updateBuildPreparationArgsForCall)
}

func (fake *FakePipelineDB) UpdateBuildPreparationArgsForCall(i int) db.BuildPreparation {
	fake.updateBuildPreparationMutex.RLock()
	defer fake.updateBuildPreparationMutex.RUnlock()
	return fake.updateBuildPreparationArgsForCall[i].prep
}

func (fake *FakePipelineDB) UpdateBuildPreparationReturns(result1 error) {
	fake.UpdateBuildPreparationStub = nil
	fake.updateBuildPreparationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePipelineDB) IsPaused() (bool, error) {
	fake.isPausedMutex.Lock()
	fake.isPausedArgsForCall = append(fake.isPausedArgsForCall, struct{}{})
	fake.isPausedMutex.Unlock()
	if fake.IsPausedStub != nil {
		return fake.IsPausedStub()
	} else {
		return fake.isPausedReturns.result1, fake.isPausedReturns.result2
	}
}

func (fake *FakePipelineDB) IsPausedCallCount() int {
	fake.isPausedMutex.RLock()
	defer fake.isPausedMutex.RUnlock()
	return len(fake.isPausedArgsForCall)
}

func (fake *FakePipelineDB) IsPausedReturns(result1 bool, result2 error) {
	fake.IsPausedStub = nil
	fake.isPausedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineDB) LoadVersionsDB() (*algorithm.VersionsDB, error) {
	fake.loadVersionsDBMutex.Lock()
	fake.loadVersionsDBArgsForCall = append(fake.loadVersionsDBArgsForCall, struct{}{})
	fake.loadVersionsDBMutex.Unlock()
	if fake.LoadVersionsDBStub != nil {
		return fake.LoadVersionsDBStub()
	} else {
		return fake.loadVersionsDBReturns.result1, fake.loadVersionsDBReturns.result2
	}
}

func (fake *FakePipelineDB) LoadVersionsDBCallCount() int {
	fake.loadVersionsDBMutex.RLock()
	defer fake.loadVersionsDBMutex.RUnlock()
	return len(fake.loadVersionsDBArgsForCall)
}

func (fake *FakePipelineDB) LoadVersionsDBReturns(result1 *algorithm.VersionsDB, result2 error) {
	fake.LoadVersionsDBStub = nil
	fake.loadVersionsDBReturns = struct {
		result1 *algorithm.VersionsDB
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineDB) GetNextInputVersions(versions *algorithm.VersionsDB, job string, inputs []config.JobInput) ([]db.BuildInput, bool, error) {
	fake.getNextInputVersionsMutex.Lock()
	fake.getNextInputVersionsArgsForCall = append(fake.getNextInputVersionsArgsForCall, struct {
		versions *algorithm.VersionsDB
		job      string
		inputs   []config.JobInput
	}{versions, job, inputs})
	fake.getNextInputVersionsMutex.Unlock()
	if fake.GetNextInputVersionsStub != nil {
		return fake.GetNextInputVersionsStub(versions, job, inputs)
	} else {
		return fake.getNextInputVersionsReturns.result1, fake.getNextInputVersionsReturns.result2, fake.getNextInputVersionsReturns.result3
	}
}

func (fake *FakePipelineDB) GetNextInputVersionsCallCount() int {
	fake.getNextInputVersionsMutex.RLock()
	defer fake.getNextInputVersionsMutex.RUnlock()
	return len(fake.getNextInputVersionsArgsForCall)
}

func (fake *FakePipelineDB) GetNextInputVersionsArgsForCall(i int) (*algorithm.VersionsDB, string, []config.JobInput) {
	fake.getNextInputVersionsMutex.RLock()
	defer fake.getNextInputVersionsMutex.RUnlock()
	return fake.getNextInputVersionsArgsForCall[i].versions, fake.getNextInputVersionsArgsForCall[i].job, fake.getNextInputVersionsArgsForCall[i].inputs
}

func (fake *FakePipelineDB) GetNextInputVersionsReturns(result1 []db.BuildInput, result2 bool, result3 error) {
	fake.GetNextInputVersionsStub = nil
	fake.getNextInputVersionsReturns = struct {
		result1 []db.BuildInput
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePipelineDB) UseInputsForBuild(buildID int, inputs []db.BuildInput) error {
	fake.useInputsForBuildMutex.Lock()
	fake.useInputsForBuildArgsForCall = append(fake.useInputsForBuildArgsForCall, struct {
		buildID int
		inputs  []db.BuildInput
	}{buildID, inputs})
	fake.useInputsForBuildMutex.Unlock()
	if fake.UseInputsForBuildStub != nil {
		return fake.UseInputsForBuildStub(buildID, inputs)
	} else {
		return fake.useInputsForBuildReturns.result1
	}
}

func (fake *FakePipelineDB) UseInputsForBuildCallCount() int {
	fake.useInputsForBuildMutex.RLock()
	defer fake.useInputsForBuildMutex.RUnlock()
	return len(fake.useInputsForBuildArgsForCall)
}

func (fake *FakePipelineDB) UseInputsForBuildArgsForCall(i int) (int, []db.BuildInput) {
	fake.useInputsForBuildMutex.RLock()
	defer fake.useInputsForBuildMutex.RUnlock()
	return fake.useInputsForBuildArgsForCall[i].buildID, fake.useInputsForBuildArgsForCall[i].inputs
}

func (fake *FakePipelineDB) UseInputsForBuildReturns(result1 error) {
	fake.UseInputsForBuildStub = nil
	fake.useInputsForBuildReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePipelineDB) CreateJobBuild(job string) (db.Build, error) {
	fake.createJobBuildMutex.Lock()
	fake.createJobBuildArgsForCall = append(fake.createJobBuildArgsForCall, struct {
		job string
	}{job})
	fake.createJobBuildMutex.Unlock()
	if fake.CreateJobBuildStub != nil {
		return fake.CreateJobBuildStub(job)
	} else {
		return fake.createJobBuildReturns.result1, fake.createJobBuildReturns.result2
	}
}

func (fake *FakePipelineDB) CreateJobBuildCallCount() int {
	fake.createJobBuildMutex.RLock()
	defer fake.createJobBuildMutex.RUnlock()
	return len(fake.createJobBuildArgsForCall)
}

func (fake *FakePipelineDB) CreateJobBuildArgsForCall(i int) string {
	fake.createJobBuildMutex.RLock()
	defer fake.createJobBuildMutex.RUnlock()
	return fake.createJobBuildArgsForCall[i].job
}

func (fake *FakePipelineDB) CreateJobBuildReturns(result1 db.Build, result2 error) {
	fake.CreateJobBuildStub = nil
	fake.createJobBuildReturns = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineDB) CreateJobBuildForCandidateInputs(job string) (db.Build, bool, error) {
	fake.createJobBuildForCandidateInputsMutex.Lock()
	fake.createJobBuildForCandidateInputsArgsForCall = append(fake.createJobBuildForCandidateInputsArgsForCall, struct {
		job string
	}{job})
	fake.createJobBuildForCandidateInputsMutex.Unlock()
	if fake.CreateJobBuildForCandidateInputsStub != nil {
		return fake.CreateJobBuildForCandidateInputsStub(job)
	} else {
		return fake.createJobBuildForCandidateInputsReturns.result1, fake.createJobBuildForCandidateInputsReturns.result2, fake.createJobBuildForCandidateInputsReturns.result3
	}
}

func (fake *FakePipelineDB) CreateJobBuildForCandidateInputsCallCount() int {
	fake.createJobBuildForCandidateInputsMutex.RLock()
	defer fake.createJobBuildForCandidateInputsMutex.RUnlock()
	return len(fake.createJobBuildForCandidateInputsArgsForCall)
}

func (fake *FakePipelineDB) CreateJobBuildForCandidateInputsArgsForCall(i int) string {
	fake.createJobBuildForCandidateInputsMutex.RLock()
	defer fake.createJobBuildForCandidateInputsMutex.RUnlock()
	return fake.createJobBuildForCandidateInputsArgsForCall[i].job
}

func (fake *FakePipelineDB) CreateJobBuildForCandidateInputsReturns(result1 db.Build, result2 bool, result3 error) {
	fake.CreateJobBuildForCandidateInputsStub = nil
	fake.createJobBuildForCandidateInputsReturns = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePipelineDB) UpdateBuildToScheduled(buildID int) (bool, error) {
	fake.updateBuildToScheduledMutex.Lock()
	fake.updateBuildToScheduledArgsForCall = append(fake.updateBuildToScheduledArgsForCall, struct {
		buildID int
	}{buildID})
	fake.updateBuildToScheduledMutex.Unlock()
	if fake.UpdateBuildToScheduledStub != nil {
		return fake.UpdateBuildToScheduledStub(buildID)
	} else {
		return fake.updateBuildToScheduledReturns.result1, fake.updateBuildToScheduledReturns.result2
	}
}

func (fake *FakePipelineDB) UpdateBuildToScheduledCallCount() int {
	fake.updateBuildToScheduledMutex.RLock()
	defer fake.updateBuildToScheduledMutex.RUnlock()
	return len(fake.updateBuildToScheduledArgsForCall)
}

func (fake *FakePipelineDB) UpdateBuildToScheduledArgsForCall(i int) int {
	fake.updateBuildToScheduledMutex.RLock()
	defer fake.updateBuildToScheduledMutex.RUnlock()
	return fake.updateBuildToScheduledArgsForCall[i].buildID
}

func (fake *FakePipelineDB) UpdateBuildToScheduledReturns(result1 bool, result2 error) {
	fake.UpdateBuildToScheduledStub = nil
	fake.updateBuildToScheduledReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineDB) GetJobBuildForInputs(job string, inputs []db.BuildInput) (db.Build, bool, error) {
	fake.getJobBuildForInputsMutex.Lock()
	fake.getJobBuildForInputsArgsForCall = append(fake.getJobBuildForInputsArgsForCall, struct {
		job    string
		inputs []db.BuildInput
	}{job, inputs})
	fake.getJobBuildForInputsMutex.Unlock()
	if fake.GetJobBuildForInputsStub != nil {
		return fake.GetJobBuildForInputsStub(job, inputs)
	} else {
		return fake.getJobBuildForInputsReturns.result1, fake.getJobBuildForInputsReturns.result2, fake.getJobBuildForInputsReturns.result3
	}
}

func (fake *FakePipelineDB) GetJobBuildForInputsCallCount() int {
	fake.getJobBuildForInputsMutex.RLock()
	defer fake.getJobBuildForInputsMutex.RUnlock()
	return len(fake.getJobBuildForInputsArgsForCall)
}

func (fake *FakePipelineDB) GetJobBuildForInputsArgsForCall(i int) (string, []db.BuildInput) {
	fake.getJobBuildForInputsMutex.RLock()
	defer fake.getJobBuildForInputsMutex.RUnlock()
	return fake.getJobBuildForInputsArgsForCall[i].job, fake.getJobBuildForInputsArgsForCall[i].inputs
}

func (fake *FakePipelineDB) GetJobBuildForInputsReturns(result1 db.Build, result2 bool, result3 error) {
	fake.GetJobBuildForInputsStub = nil
	fake.getJobBuildForInputsReturns = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePipelineDB) GetNextPendingBuild(job string) (db.Build, bool, error) {
	fake.getNextPendingBuildMutex.Lock()
	fake.getNextPendingBuildArgsForCall = append(fake.getNextPendingBuildArgsForCall, struct {
		job string
	}{job})
	fake.getNextPendingBuildMutex.Unlock()
	if fake.GetNextPendingBuildStub != nil {
		return fake.GetNextPendingBuildStub(job)
	} else {
		return fake.getNextPendingBuildReturns.result1, fake.getNextPendingBuildReturns.result2, fake.getNextPendingBuildReturns.result3
	}
}

func (fake *FakePipelineDB) GetNextPendingBuildCallCount() int {
	fake.getNextPendingBuildMutex.RLock()
	defer fake.getNextPendingBuildMutex.RUnlock()
	return len(fake.getNextPendingBuildArgsForCall)
}

func (fake *FakePipelineDB) GetNextPendingBuildArgsForCall(i int) string {
	fake.getNextPendingBuildMutex.RLock()
	defer fake.getNextPendingBuildMutex.RUnlock()
	return fake.getNextPendingBuildArgsForCall[i].job
}

func (fake *FakePipelineDB) GetNextPendingBuildReturns(result1 db.Build, result2 bool, result3 error) {
	fake.GetNextPendingBuildStub = nil
	fake.getNextPendingBuildReturns = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePipelineDB) SaveResourceVersions(arg1 atc.ResourceConfig, arg2 []atc.Version) error {
	fake.saveResourceVersionsMutex.Lock()
	fake.saveResourceVersionsArgsForCall = append(fake.saveResourceVersionsArgsForCall, struct {
		arg1 atc.ResourceConfig
		arg2 []atc.Version
	}{arg1, arg2})
	fake.saveResourceVersionsMutex.Unlock()
	if fake.SaveResourceVersionsStub != nil {
		return fake.SaveResourceVersionsStub(arg1, arg2)
	} else {
		return fake.saveResourceVersionsReturns.result1
	}
}

func (fake *FakePipelineDB) SaveResourceVersionsCallCount() int {
	fake.saveResourceVersionsMutex.RLock()
	defer fake.saveResourceVersionsMutex.RUnlock()
	return len(fake.saveResourceVersionsArgsForCall)
}

func (fake *FakePipelineDB) SaveResourceVersionsArgsForCall(i int) (atc.ResourceConfig, []atc.Version) {
	fake.saveResourceVersionsMutex.RLock()
	defer fake.saveResourceVersionsMutex.RUnlock()
	return fake.saveResourceVersionsArgsForCall[i].arg1, fake.saveResourceVersionsArgsForCall[i].arg2
}

func (fake *FakePipelineDB) SaveResourceVersionsReturns(result1 error) {
	fake.SaveResourceVersionsStub = nil
	fake.saveResourceVersionsReturns = struct {
		result1 error
	}{result1}
}

var _ scheduler.PipelineDB = new(FakePipelineDB)
