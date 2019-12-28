package youtrack

import (
	"bytes"
	"context"
	"net/url"
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

	issueBody := `
This is a test issue that uses *some* markdown.

### This is a Heading 3
This is a [link](https://example.com).

Here is a code snippet:
` + "```" + `
class Test {

};
` + "```"

	issue, err := api.CreateIssue(ctx, projectID, "Test Issue", issueBody)
	if err != nil {
		t.Fatal("Failed to create test issue.", err)
	}
	if issue.ID == "" {
		t.Fatal("Empty Issue ID.")
	}
	if issue.NumberInProject == 0 {
		t.Fatal("NumberInProject is 0")
	}

	attachmentData := bytes.NewReader([]byte(`Test Attachment Data`))
	attachmentID, err := api.CreateIssueAttachment(ctx, issue.ID, attachmentData, "myAttachment.txt", "text/plain")
	if err != nil {
		t.Fatal("Failed to create issue attachment.", err)
	}
	if attachmentID == "" {
		t.Fatal("Empty attachment ID")
	}

	u := api.IssueURL("TP", issue.NumberInProject)
	if u.String() == "" {
		t.Error("Expected u.String() to not be empty.")
	}
}

func TestIssuesURL(t *testing.T) {
	base, err := url.Parse("https://example.com/youtrack/api/")
	if err != nil {
		t.Fatal("Failed to parse example base URL.", err)
	}
	u := IssueURL(base, "TP", 123)
	expected := "https://example.com/youtrack/issue/TP-123"
	if u.String() != expected {
		t.Fatal(u.String(), "is not expected value", expected)
	}
}
