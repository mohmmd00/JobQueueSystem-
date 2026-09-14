// Manager
package jobaggregate

import (
	"errors"
	"sort"

	"github.com/google/uuid"
)

type Manager struct {
	jobs map[uuid.UUID]*Job //jobs is a map where each UUID points! to a Job.
}

func NewManager() Manager { //constructor
	return Manager{jobs: make(map[uuid.UUID]*Job)}
}

func (m *Manager) CreateJob(name string, priority int) (*Job, error) {

	job := NewJob(name, priority)
	m.jobs[job.ID] = &job //gives the map the memory address of that new job to keep

	return &job, nil

}
func (m *Manager) ListJobs() ([]*Job, error) {

	if len(m.jobs) == 0 {
		return nil, errors.New("Couldnt find any Jobs. The map is empty.")
	}

	var jobsHolder []*Job
	for _, job := range m.jobs {
		jobsHolder = append(jobsHolder, job)
	}
	return jobsHolder, nil
}
func (m *Manager) FindJob(id string) (*Job, error) {

	parsedUUID, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}
	foundJob, ok := m.jobs[parsedUUID]
	if !ok {
		return nil, errors.New("Job finding failed. couldnt find job")
	}
	return foundJob, nil

}
func (m *Manager) CancelJob(id string) error {

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	foundJob, ok := m.jobs[parsedUUID]
	if !ok {
		return errors.New("Job canceling failed. couldnt find job.")
	}
	if foundJob.Status != Pending {
		return errors.New("Founded job is not on pending.")
	}
	foundJob.Status = Cancelled // job status changed to cancel
	return nil

}
func (m *Manager) pendingJobs() ([]*Job, error) {

	if len(m.jobs) == 0 {
		return nil, errors.New("Couldnt find any Jobs. The map is empty.")
	}

	var jobsHolder []*Job

	for _, job := range m.jobs { // for each job in map(jobs)
		if job.Status == Pending {
			jobsHolder = append(jobsHolder, job) //pick up the jobs on pending into holder

		}
	}
	if len(jobsHolder) == 0 {
		return nil, errors.New("Couldnt find any Pending Jobs.")
	}

	sort.Slice(jobsHolder, func(i, j int) bool {
		return jobsHolder[i].Priority > jobsHolder[j].Priority
	})

	return jobsHolder, nil

}

/*
func (m *Manager) processJob(job *Job) {

		job.Status = Running

		time.Sleep(2 * time.Second)

		job.Status = Completed
	}
*/
/*func (m *Manager) ProcessQueue() error {

	jobs, err := m.pendingJobs()
	if err != nil {
		return err
	}
	for _, job := range jobs {
	}
	return nil
}
*/
func (m *Manager) Statistics() (*JobStatistics, error) {

	var statistics JobStatistics

	if len(m.jobs) == 0 {
		return nil, errors.New("Couldnt find any Jobs. The map is empty.")
	}

	for _, job := range m.jobs {
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

	return &statistics, nil
}
