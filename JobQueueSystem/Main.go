// Main
package main

import (
	JobAgg "JobQueueSystem/JobQueueSystem/JobAggregate"
	natsss "JobQueueSystem/JobQueueSystem/Messaging"
	Worker "JobQueueSystem/JobQueueSystem/Worker"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nats-io/nats.go"
)

const menu = `
=== Job Queue Manager ===
1. Create Job
2. List Jobs
3. Find Job
4. Cancel Job
5. Process Queue
6. Show Statistics
7. Exit`

const jobHeader = "ID                                   Name                 Priority    Status       Created At"

var stdin = bufio.NewReader(os.Stdin)

// ask prints a label and returns the trimmed line the user typed.
func ask(label string) string {
	fmt.Print(label)
	line, err := stdin.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(line)
}

func printJob(job JobAgg.Job) {
	fmt.Printf("%v %v %v %v %v \n", job.ID, job.Name, job.Priority, job.Status, job.CreatedAt)
}

func main() {

	nc, connErr := natsss.Connect()
	if connErr != nil {
		log.Fatal(connErr)
	}
	defer nc.Close()

	if workerErr := Worker.Start(nc); workerErr != nil {
		log.Fatal(workerErr)
	}

	Manager := JobAgg.NewManager() // make new empty map by calling constructor of job manager

	_, subErr := nc.Subscribe("jobs.completed", func(msg *nats.Msg) { // subscribe to jobs.completed subject
		var subJob JobAgg.Job

		if unmarshErr := json.Unmarshal(msg.Data, &subJob); unmarshErr != nil {
			log.Println("Failed to decode jobs:", unmarshErr)
			return
		}
		if err := Manager.UpdateJobStatus(subJob.ID.String(), subJob.Status); err != nil {
			log.Println("Failed to update: ", err)
		}
	})
	if subErr != nil {
		log.Println("Failed to subscribe: ", subErr)
	}

	for {
		fmt.Println(menu)
		choice := ask("Enter your choice: ")

		switch choice {

		case "1":
			fmt.Println("Create Job selected.")

			name := ask("Enter job name: ")
			if name == "" {
				fmt.Println("Job name cannot be empty.")
				continue
			}

			priority, convErr := strconv.Atoi(ask("Enter priority: "))
			if convErr != nil {
				fmt.Println("Priority must be a valid number.")
				continue
			}

			createdJob := Manager.CreateJob(name, priority)
			fmt.Println("Job created successfully.")
			fmt.Println(createdJob.ID)

		case "2":
			fmt.Println("List Jobs selected.")

			fetchedJobs, listErr := Manager.ListJobs()
			if listErr != nil {
				fmt.Println("Failed to list jobs.")
				fmt.Println(listErr.Error())
				continue
			}

			fmt.Println("Jobs listed successfully.")
			fmt.Println(jobHeader)
			for _, job := range fetchedJobs {
				printJob(*job)
			}

		case "3":
			fmt.Println("Find Job selected.")

			id := ask("Enter job UUID: ")

			job, findErr := Manager.FindJob(id)
			if findErr != nil {
				fmt.Println("Job not found.")
				fmt.Println(findErr.Error())
				continue
			}

			fmt.Println("Job found.")
			fmt.Println(jobHeader)
			printJob(*job)

		case "4":
			fmt.Println("Cancel Job selected.")

			id := ask("Enter job UUID: ")

			if cancelErr := Manager.CancelJob(id); cancelErr != nil {
				fmt.Println("Job cannot be cancelled.")
				fmt.Println(cancelErr.Error())
				continue
			}

			fmt.Println("Job cancelled successfully.")

		case "5":
			fmt.Println("Process Queue selected.")

			jobs, pendingErr := Manager.PendingJobs()
			if pendingErr != nil {
				log.Println("Failed to gather jobs: ", pendingErr)
				fmt.Println("No pending jobs in the queue.")
				continue
			}

			failed := false
			for _, job := range jobs {
				data, marshErr := json.Marshal(job)
				if marshErr != nil {
					log.Println("Failed to marshal jobs: ", marshErr)
					continue
				}
				if pubErr := nc.Publish("jobs.process", data); pubErr != nil {
					log.Println("Failed to publish: ", pubErr)
					fmt.Println("No pending jobs in the queue.")
					failed = true
					break
				}
			}
			if failed {
				continue
			}

			fmt.Println("Queue processing finished.")

		case "6":
			fmt.Println("Show Statistics selected.")

			stats, statsErr := Manager.Statistics()
			if statsErr != nil {
				fmt.Println("No jobs available.")
				fmt.Println(statsErr.Error())
				continue
			}

			fmt.Println("=== Job Statistics ===")
			fmt.Printf("%+v\n", stats)

		case "7":
			fmt.Println("Exit selected.")
			fmt.Println("Goodbye.")
			return

		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}