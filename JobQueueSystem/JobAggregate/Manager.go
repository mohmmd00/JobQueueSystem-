// Manager
package jobaggregate

import (
	"errors"
	"sort"
	"sync"

	"github.com/google/uuid"
)

type Manager struct {
	jobs map[uuid.UUID]*Job //jobs is a map where each UUID points! to a Job.
	mu   sync.Mutex
}

func NewManager() *Manager { //constructor
	return &Manager{jobs: make(map[uuid.UUID]*Job)}

}

func (m *Manager) CreateJob(name string, priority int) *Job {
	m.mu.Lock()
	defer m.mu.Unlock()

	job := NewJob(name, priority)
	m.jobs[job.ID] = &job //gives the map the memory address of that new job to keep

	return &job

}
func (m *Manager) ListJobs() ([]*Job, error) {

	m.mu.Lock()
	defer m.mu.Unlock()

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
	m.mu.Lock()
	defer m.mu.Unlock()

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
	m.mu.Lock()
	defer m.mu.Unlock()

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
func (m *Manager) PendingJobs() ([]Job, error) {

	var jobsHolder []Job
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.jobs) == 0 {
		return nil, errors.New("Couldnt find any Jobs. The map is empty.")
	}

	for _, job := range m.jobs { // for each job in map(jobs)
		if job.Status == Pending {
			jobsHolder = append(jobsHolder, *job) //pick up the jobs on pending into holder

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

func (m *Manager) Statistics() (*JobStatistics, error) {

	var statistics JobStatistics
	m.mu.Lock()
	defer m.mu.Unlock()

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

func (m *Manager) UpdateJobStatus(id string, status JobStatus) error {

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	foundJob, ok := m.jobs[parsedUUID]
	if !ok {
		return errors.New("Job status updating failed. couldnt find job.")
	}

	foundJob.Status = status

	return nil

}
