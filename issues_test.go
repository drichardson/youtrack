package youtrack

import (
	"context"
	"testing"
	"time"
)

func TestIssuesSystem(t *testing.T) {

	api, err := NewDefaultApi()
	if err != nil {
		t.Fatal("Failed to create api.", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	projectID, err := api.ProjectIDForShortName(ctx, "TP")
	if err != nil {
		t.Fatal("Failed to lookup project ID for TP", err)
	}

	err = api.CreateIssue(ctx, projectID, "Test Issue", "Test Issue Body")
	if err != nil {
		t.Fatal("Failed to create test issue.", err)
	}
}
