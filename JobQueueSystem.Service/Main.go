// Main
package main

import (
	JobAgg "JobQueueSystem/JobQueueSystem.Service/JobAggregate"
	MessageHolder "JobQueueSystem/JobQueueSystem.Service/MessageAggregate"
	"fmt"
)

func main() {
	Manager := JobAgg.NewManager() // make new empty map by calling constructor of job manager

	for {
		MessageHolder.ShowOptions()

		var optionInputHolder int
		fmt.Scanln(&optionInputHolder)

		switch optionInputHolder {
		case 1:

			MessageHolder.ShowSelected("CreateJob")
			MessageHolder.ShowEnterJobInformation()

			var jobNameHolder string
			var jobPriority int
			fmt.Scanln(&jobNameHolder)
			fmt.Scanln(&jobPriority)

			FetchedCreatedJob, err := Manager.CreateJob(jobNameHolder, jobPriority)
			if err != nil {
				MessageHolder.ShowFailed("CreateJob")
				println(err)

			}
			MessageHolder.ShowSuccessful("CreateJob")
			fmt.Println(FetchedCreatedJob.ID)

		case 2:

			MessageHolder.ShowSelected("ListJobs")
			fetchedJobs, err := Manager.ListJobs()

			if err != nil {
				MessageHolder.ShowFailed("ListJobs")
				println(err)
			}
			if len(fetchedJobs) == 0 {
				MessageHolder.ShowFailed("FindJob") // job not found but result of that operation was successful !
			}
			MessageHolder.ShowSuccessful("ListJobs")
			MessageHolder.ShowJobDetails()
			for _, job := range fetchedJobs {
				fmt.Printf("%v %v %v %v %v \n", job.ID, job.Name, job.Priority, job.Status, job.CreatedAt)
			}

		case 3:

			MessageHolder.ShowSelected("FindJob")
			MessageHolder.ShowEnterJobId()

			var JobIdHolder string
			fmt.Scanln(&JobIdHolder)

			job, err := Manager.FindJob(JobIdHolder)
			if err != nil {
				MessageHolder.ShowFailed("FindJob")
			}
			MessageHolder.ShowSuccessful("FindJob")
			MessageHolder.ShowJobDetails()
			fmt.Printf("%v %v %v %v %v \n", job.ID, job.Name, job.Priority, job.Status, job.CreatedAt)

		case 4:

			MessageHolder.ShowSelected("CancelJob")
			MessageHolder.ShowEnterJobId()

			var JobIdHolder string
			fmt.Scanln(&JobIdHolder)

			err := Manager.CancelJob(JobIdHolder)
			if err != nil {
				MessageHolder.ShowFailed("CancelJob")
			}
			MessageHolder.ShowSuccessful("CancelJob")

		case 5:

			MessageHolder.ShowSelected("ProcessQueue")
			err := Manager.ProcessQueue()
			if err != nil {
				MessageHolder.ShowFailed("ProcessQueue")
			}
			MessageHolder.ShowSuccessful("ProcessQueue")

		case 6:
			MessageHolder.ShowSelected("Statistics")

			StatisticHolder, err := Manager.Statistics()
			if err != nil {
				MessageHolder.ShowFailed("Statistics")
			}
			MessageHolder.ShowSuccessful("Statistics")
			fmt.Printf("%+v\n", StatisticHolder)

		case 7:
			MessageHolder.ShowSelected("Exit")
			MessageHolder.ShowSuccessful("Exit")
			return

		default:
			MessageHolder.ShowInvalidChoice()

		}
	}

}
