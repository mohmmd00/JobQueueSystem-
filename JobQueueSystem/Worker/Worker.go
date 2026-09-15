package worker

import (
	JobAgg "JobQueueSystem/JobQueueSystem/JobAggregate"
	NatsAgg "JobQueueSystem/JobQueueSystem/Messaging"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func Start(nc *nats.Conn) error {
	_, subErr := NatsAgg.Subscribe(nc, "jobs.process", func(msg *nats.Msg) {
		var job JobAgg.Job

		unmarshErr := json.Unmarshal(msg.Data, &job)
		if unmarshErr != nil {
			log.Println("Failed to decode jobs:", unmarshErr)
			return
		}

		newJob, marshErr := json.Marshal(processJob(job))

		if marshErr != nil {
			log.Println("Failed to decode jobs:", marshErr)
			return
		}

		pubErr := NatsAgg.Publish(nc, "jobs.completed", newJob)
		if pubErr != nil {
			log.Println("Failed to Publish complete Job:", pubErr)

		}

	})

	if subErr != nil {
		nc.Close()
		return subErr
	}

	fmt.Println("Worker is listening...")
	return nil
}

func processJob(job JobAgg.Job) *JobAgg.Job {

		fmt.Println("\nProcessing:", job.Name)

		job.Status = JobAgg.Running

		time.Sleep(3 * time.Second)

		job.Status = JobAgg.Completed

		fmt.Println("\nCompleted:", job.Name)
	return &job
}
