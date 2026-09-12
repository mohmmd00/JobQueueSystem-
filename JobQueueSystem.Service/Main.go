// Main
package main

import (
	Jobee "JobQueueSystem/JobQueueSystem.Service/JobAggregate"
	MessageHolder "JobQueueSystem/JobQueueSystem.Service/MessageAggregate"
	"fmt"
)

func main() {
	Manager := Jobee.NewJobManager() // make new empty map by calling constructor of job manager

	for {
		MessageHolder.ShowOptions()

		var OptionInputHolder int
		fmt.Scanln(&OptionInputHolder)

		switch OptionInputHolder {
		case 1:

			MessageHolder.ShowSelected("CreateJob")
			MessageHolder.ShowEnterJobInformation()

			var jobNameHolder string
			var jobPriority int
			fmt.Scanln(&jobNameHolder)
			fmt.Scanln(&jobPriority)

			FetchedCreatedJob, Result := Manager.CreateJob(jobNameHolder, jobPriority)
			if !Result {
				MessageHolder.ShowFailed("CreateJob")

			} else {
				MessageHolder.ShowSuccessful("CreateJob")
				fmt.Println(FetchedCreatedJob.ID)
			}

		case 2:

			MessageHolder.ShowSelected("ListJobs")
			jobAsSlice, result := Manager.ListJob()

			if !result {
				MessageHolder.ShowFailed("ListJobs")
			}
			if len(jobAsSlice) == 0 {
				MessageHolder.ShowFailed("FindJob") // job not found but result of that operation was successful !
			} else {
				MessageHolder.ShowSuccessful("ListJobs")
				MessageHolder.ShowJobDetails()
				for _, job := range jobAsSlice {
					fmt.Printf("%v %v %v %v %v \n", job.ID, job.Name, job.Priority, job.Status, job.CreatedAt)
				}
			}

		case 3:

			MessageHolder.ShowSelected("FindJob")
			MessageHolder.ShowEnterJobId()

			var JobIdHolder string
			fmt.Scanln(&JobIdHolder)

			job, Result := Manager.FindJob(JobIdHolder)
			if !Result {
				MessageHolder.ShowFailed("FindJob")
			} else {
				MessageHolder.ShowSuccessful("FindJob")
				MessageHolder.ShowJobDetails()
				fmt.Printf("%v %v %v %v %v \n", job.ID, job.Name, job.Priority, job.Status, job.CreatedAt)
			}

		case 4:

			MessageHolder.ShowSelected("CancelJob")
			MessageHolder.ShowEnterJobId()

			var JobIdHolder string
			fmt.Scanln(&JobIdHolder)

			Result := Manager.CancelJob(JobIdHolder)
			if !Result {
				MessageHolder.ShowFailed("CancelJob")
			} else {
				MessageHolder.ShowSuccessful("CancelJob")
			}

		case 5:

			MessageHolder.ShowSelected("ProcessQueue")
			Result := Manager.ProcessQueue()
			if !Result {
				MessageHolder.ShowFailed("ProcessQueue")
			} else {
				MessageHolder.ShowSuccessful("ProcessQueue")
			}

		case 6:
			MessageHolder.ShowSelected("Statistics")

			StatisticHolder, Result := Manager.Statistics()
			if !Result {
				MessageHolder.ShowFailed("Statistics")
			} else {
				MessageHolder.ShowSuccessful("Statistics")
				println(StatisticHolder)
			}
		case 7:
			MessageHolder.ShowSelected("Exit")
			MessageHolder.ShowSuccessful("Exit")
			return

		default:
			MessageHolder.ShowInvalidChoice()

		}
	}

}
