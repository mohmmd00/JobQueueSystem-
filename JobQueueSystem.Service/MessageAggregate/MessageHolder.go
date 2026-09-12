package MessageAggregate

import "fmt"

type ActionMessages struct{ Selected, Successful, Failed string }

var Actions = map[string]ActionMessages{
	"CreateJob":       {"Create Job selected.", "Job created successfully.", "Failed to create job."},
	"ListJobs":        {"List Jobs selected.", "Jobs listed successfully.", "Failed to list jobs."},
	"FindJob":         {"Find Job selected.", "Job found.", "Job not found."},
	"CancelJob":       {"Cancel Job selected.", "Job cancelled successfully.", "Job cannot be cancelled."},
	"ChangeJobStatus": {"Change Job Status selected.", "Job status changed successfully.", "Invalid job status."},
	"ProcessQueue":    {"Process Queue selected.", "Queue processing finished.", "No pending jobs in the queue."},
	"ProcessJob":      {"Job processing started.", "Job completed successfully.", "Job processing failed."},
	"Statistics":      {"Show Statistics selected.", "=== Job Statistics ===", "No jobs available."},
	"Exit":            {"Exit selected.", "Goodbye.", ""},
}

const (
	MsgJobDetail       = "ID                                   Name                 Priority    Status       Created At"
	MsgInvalidChoice   = "Invalid choice. Please select a valid option."
	MsgInvalidUUID     = "Invalid UUID."
	MsgInvalidName     = "Job name cannot be empty."
	MsgInvalidPriority = "Priority must be a valid number."

	MsgMenuTitle = "=== Job Queue Manager ===\n1. Create Job\n2. List Jobs\n3. Find Job\n4. Cancel Job\n5. Process Queue\n6. Show Statistics\n7. Exit"

	MsgEnterChoice   = "Enter your choice: "
	MsgEnterJobName  = "Enter job name: "
	MsgEnterPriority = "Enter priority: "
	MsgEnterJobID    = "Enter job UUID: "
)

func ShowSelected(action string)   { fmt.Println(Actions[action].Selected) }
func ShowSuccessful(action string) { fmt.Println(Actions[action].Successful) }
func ShowFailed(action string)     { fmt.Println(Actions[action].Failed) }

func ShowOptions()             { fmt.Println(MsgMenuTitle); fmt.Print(MsgEnterChoice) }
func ShowJobDetails()          { fmt.Println(MsgJobDetail) }
func ShowInvalidChoice()       { fmt.Println(MsgInvalidChoice) }
func ShowEnterJobInformation() { fmt.Println(MsgEnterJobName + "\n" + MsgEnterPriority) }
func ShowEnterJobId()          { fmt.Println(MsgEnterJobID) }
