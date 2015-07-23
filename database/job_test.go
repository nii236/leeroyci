package database

import (
	"testing"
)

func TestCGDoneJob(t *testing.T) {
	repo, _ := CreateRepository("foo", "baz", "accessKey", false, false, false)

	job := CreateJob(repo, "branch", "commit", "commitURL", "name", "email")
	job.TasksDone()
	job.DeployDone()
	get := GetJob(job.ID)

	if job.TasksFinished == get.TasksFinished {
		t.Error("tasks not finished")
	}

	if job.DeployFinished == get.DeployFinished {
		t.Error("deploy not finished")
	}
}
