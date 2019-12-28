package youtrack

import (
	"bytes"
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

	issueID, err := api.CreateIssue(ctx, projectID, "Test Issue", "Test Issue Body")
	if err != nil {
		t.Fatal("Failed to create test issue.", err)
	}
	if issueID == "" {
		t.Fatal("Empty Issue ID.")
	}

	attachmentData := bytes.NewReader([]byte(`Test Attachment Data`))
	attachmentID, err := api.CreateIssueAttachment(ctx, issueID, attachmentData, "myAttachment.txt", "text/plain")
	if err != nil {
		t.Fatal("Failed to create issue attachment.", err)
	}
	if attachmentID == "" {
		t.Fatal("Empty attachment ID")
	}
}
