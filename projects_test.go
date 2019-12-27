package youtrack

import (
	"context"
	"testing"
	"time"
)

func TestProjectsSystem(t *testing.T) {

	api, err := NewDefaultApi()
	if err != nil {
		t.Fatal("Failed to create api.", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	id, err := api.ProjectIDForShortName(ctx, "TP")
	if err != nil {
		t.Fatal("ProjectIDForShortName failed.", err)
	}
	if id == "" {
		t.Fatal("ProjectIDForShortName returned empty id.")
	}
}
