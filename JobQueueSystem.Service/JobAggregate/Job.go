// job
package jobaggregate

import (
	"time"

	"github.com/google/uuid"
)

type JobStatus string

const (
	Pending   JobStatus = "pending"
	Running   JobStatus = "running"
	Completed JobStatus = "completed"
	Failed    JobStatus = "failed"
	Cancelled JobStatus = "cancelled"
)

type JobStatistics struct {
	Total     int
	Pending   int
	Running   int
	Completed int
	Failed    int
	Cancelled int
}

type Job struct {
	ID        uuid.UUID
	Name      string
	Priority  int
	Status    JobStatus
	CreatedAt time.Time
}

func NewJob(name string, priority int) Job {
	return Job{
		ID:        uuid.New(),
		Name:      name,
		Priority:  priority,
		Status:    Pending,
		CreatedAt: time.Now(),
	}
}
