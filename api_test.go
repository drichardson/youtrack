package youtrack

import (
	"context"
	"testing"
	"time"
)

func TestAPISystem(t *testing.T) {

	api, err := NewDefaultApi()
	if err != nil {
		t.Error("Failed to create api.", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = api.CreateIssueRequest(ctx, "Test Project", "Test Issue", "Test Issue Body")
	if err != nil {
		t.Error("Failed to create test issue.", err)
	}
}
