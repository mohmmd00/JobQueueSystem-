// Manager
package jobaggregate

import (
	"sort"

	"github.com/google/uuid"
)

type JobManager struct {
	Jobs map[uuid.UUID]*Job //jobs is a map where each UUID points! to a Job.
}

func NewJobManager() JobManager { //constructor
	return JobManager{Jobs: make(map[uuid.UUID]*Job)}
}

func (m *JobManager) CreateJob(name string, priority int) (createdJob Job , isSuccessful bool) {

	JobCreated := NewJob(name, priority)
	m.Jobs[JobCreated.ID] = &JobCreated

	return JobCreated, true

}
func (m *JobManager) ListJob() (jobs []*Job, isSuccessful bool) {

	var fetchedJobsAsSlice []*Job
	for _ , job := range m.Jobs{
		fetchedJobsAsSlice = append(fetchedJobsAsSlice, job)
	}
	return fetchedJobsAsSlice, true
}
func (m *JobManager) FindJob(id string) (job *Job, isSuccessful bool) {

	parsedUuid, err := uuid.Parse(id)

	if err != nil {
		return nil, false
	} else {
		FoundJob, IsExists := m.Jobs[parsedUuid]
		if !IsExists {
			return nil, false
		}
		return FoundJob, true
	}

}
func (m *JobManager) CancelJob(id string) (isSuccessful bool) {

	parsedUuid, err := uuid.Parse(id)
	if err != nil {
		return false
	}
	FoundJob, IsExists := m.Jobs[parsedUuid]
	if !IsExists {
		return false
	}
	if FoundJob.Status != Pending {
		return false
	}
	FoundJob.Status = Cancelled
	return true

}
func (m *JobManager) getPendingJobsSortedByPriority() (pendingJobsAsSlice []*Job, isSuccessful bool) {

	var pendingJobs []*Job

	for _, job := range m.Jobs {
		if job.Status == Pending {
			pendingJobs = append(pendingJobs, job)

		}
	}
	sort.Slice(pendingJobs, func(i, j int) bool {
		return pendingJobs[i].Priority > pendingJobs[j].Priority
	})

	if len(pendingJobs) == 0 {
		return nil, false
	}

	return pendingJobs, true

}
func (m *JobManager) ProcessQueue() (isSuccessful bool) {

	pendingJobs, isSuccessful := m.getPendingJobsSortedByPriority()
	if !isSuccessful {
		return false
	}
	for _, job := range pendingJobs {

		job.Status = Running

		// some bullshit we dont care

		job.Status = Completed

	}
	return true
}
func (m *JobManager) Statistics() (statistic JobStatistics, isSuccessful bool) {

	var statistics JobStatistics

	if len(m.Jobs) == 0 {
		return JobStatistics{}, false
	}

	for _, job := range m.Jobs {
		statistics.Total++

		switch job.Status {
		case Pending:
			statistics.Pending++
		case Running:
			statistics.Running++
		case Completed:
			statistics.Completed++
		case Failed:
			statistics.Failed++
		case Cancelled:
			statistics.Cancelled++
		}
	}

	return statistics, true
}
